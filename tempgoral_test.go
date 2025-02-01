package temporal

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetTipInfo_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tip_floor" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TipInfo{
			Time:                     "2023-01-01T00:00:00Z",
			LandedTips25ThPercentile: "100",
			LandedTips50ThPercentile: "200",
			LandedTips75ThPercentile: "300",
			LandedTips95ThPercentile: "400",
			LandedTips99ThPercentile: "500",
		})
	}))
	defer ts.Close()

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, network, ts.Listener.Addr().String())
			},
		},
	}

	info, err := GetTipInfo(client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := &TipInfo{
		Time:                     "2023-01-01T00:00:00Z",
		LandedTips25ThPercentile: "100",
		LandedTips50ThPercentile: "200",
		LandedTips75ThPercentile: "300",
		LandedTips95ThPercentile: "400",
		LandedTips99ThPercentile: "500",
	}
	if !reflect.DeepEqual(info, expected) {
		t.Errorf("expected %+v, got %+v", expected, info)
	}
}

func TestSubscribeTipStream_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("upgrade error: %v", err)
		}
		defer conn.Close()

		msg := []*TipInfo{{
			Time:                     "2023-01-01T00:00:00Z",
			LandedTips25ThPercentile: "100",
			LandedTips50ThPercentile: "200",
			LandedTips75ThPercentile: "300",
			LandedTips95ThPercentile: "400",
			LandedTips99ThPercentile: "500",
		}}
		if err := conn.WriteJSON(msg); err != nil {
			t.Fatalf("write error: %v", err)
		}
	}))
	defer srv.Close()

	oldURL := tipStreamURL
	tipStreamURL = "ws://" + srv.Listener.Addr().String()
	defer func() { tipStreamURL = oldURL }()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, chErr, err := SubscribeTipStream(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	select {
	case tips := <-ch:
		expected := []*TipInfo{{
			Time:                     "2023-01-01T00:00:00Z",
			LandedTips25ThPercentile: "100",
			LandedTips50ThPercentile: "200",
			LandedTips75ThPercentile: "300",
			LandedTips95ThPercentile: "400",
			LandedTips99ThPercentile: "500",
		}}
		if !reflect.DeepEqual(tips, expected) {
			t.Errorf("expected %+v, got %+v", expected, tips)
		}
	case err := <-chErr:
		t.Fatalf("unexpected error: %v", err)
	case <-ctx.Done():
		t.Fatal("timed out waiting for message")
	}
}

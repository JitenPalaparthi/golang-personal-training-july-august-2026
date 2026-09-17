# Go Live Scheduler Visualizer

Run:
    go run .

Then open: http://localhost:8080

Click **Start real goroutines**.

This launches actual Go goroutines doing CPU work, timer waits, and real HTTP I/O. Live events are streamed to the browser with Server-Sent Events.

Important accuracy note: the goroutines and HTTP I/O are real. The displayed P/M identities and per-P local queues are an educational projection. Go does not expose every internal scheduler queue as a stable live API. For a deeper runtime-accurate implementation, capture and parse `runtime/trace` scheduler events and reconstruct state transitions.

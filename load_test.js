import http from "k6/http";

const payloadURL = "https://go-million.chabk.ir/payload"
let data = {"text": "Mehrnoosh"}

export const options = {
    stages: [
        { duration: '10s', target: 10 },
        { duration: '10s', target: 200 },
        { duration: '20s', target: 200 },
        { duration: '20s', target: 0 },
    ],
    noConnectionReuse: true,
}

export default function () {
    http.post(payloadURL, JSON.stringify(data));
}

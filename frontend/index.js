const express = require("express");

const app = express();

const PORT = 3000;

app.get("/", (req, res) => {
    res.json({
        message: "Hello from Frontend",
        service: "frontend",
        status: "UP"
    });
});

app.get("/health", (req, res) => {
    res.json({
        service: "frontend",
        status: "UP"
    });
});

app.listen(PORT, "0.0.0.0", () => {
    console.log(`Frontend running on port ${PORT}`);
});
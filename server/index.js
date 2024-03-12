const http = require("http");
const express = require("express")

const app = express();

app.get("", (req, res) => {
  console.log("something")
  res.send("chalra hu");
})

app.listen(3000, () => console.log("server started on 3000"))

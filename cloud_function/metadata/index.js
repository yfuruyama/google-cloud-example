const client = require("request");
 
exports.main = (req, res) => {
  const path = req.body.path;
  const query = req.body.query;
  client.get({
    url: "http://metadata.google.internal" + path,
    headers: {
      "Metadata-Flavor": "Google",
    },
    qs: query,
  }, function (error, response, body) {
    console.log(body);
  });
  res.send("OK");
};

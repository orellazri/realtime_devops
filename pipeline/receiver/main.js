const amqp = require("amqplib/callback_api");
const moment = require("moment");

amqp.connect("amqp://guest:guest@localhost:5672", function (error0, connection) {
  if (error0) {
    throw error0;
  }

  connection.createChannel(function (error1, channel) {
    if (error1) {
      throw error1;
    }
    var exchange = "compute";

    channel.assertExchange(exchange, "fanout", {
      durable: false,
    });

    channel.assertQueue("", { exclusive: true }, function (error2, q) {
      if (error2) {
        throw error2;
      }
      q.queue = "compute";
      console.log("Waiting for messages in %s. To exit press CTRL+C", q.queue);
      channel.bindQueue(q.queue, exchange, "");

      channel.consume(
        q.queue,
        function (msg) {
          if (msg.content) {
            let msgContent = msg.content.toString().split(",");
            let msgTime = msgContent[0];
            let x = msgContent[1];
            let y = msgContent[2];
            console.log("[RabbitMQ] %s,%s", x, y);
            msgTime = moment(msgTime).add(2,"hours");
            let now = moment();
            let difference = moment.duration(now.diff(msgTime)).asMilliseconds();
            console.log(`\t${difference}ms`);
          }
        },
        { noAck: false }
      );
    });
  });
});

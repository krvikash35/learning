package kafka.getting.started;

import java.util.Properties;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;

public class MyProducer {

    private Producer<String, String> producer;

    public MyProducer(String broker) {
        Properties props = new Properties();
        props.put("bootstrap.servers", broker);
        props.put("partitioner.ignore.keys", "true");
        this.producer = new KafkaProducer<>(props, new StringSerializer(), new StringSerializer());
    }

    

    public void send(String message) {
        try {
            this.producer.send(new ProducerRecord<>("test-topic",  message));
        } catch (Exception e) {
            producer.close();
            System.out.println("producer couldn't send the message due to error: "+e);
        }
      
    }

}

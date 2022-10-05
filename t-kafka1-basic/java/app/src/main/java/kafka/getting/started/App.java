/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package kafka.getting.started;

import java.util.Arrays;
import java.util.Properties;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.errors.AuthorizationException;
import org.apache.kafka.common.errors.OutOfOrderSequenceException;
import org.apache.kafka.common.errors.ProducerFencedException;
import org.apache.kafka.common.serialization.StringSerializer;

public class App {
    public String getGreeting() {
        return "Hello World!";
    }

    public static void main(String[] args) {
        System.out.println(new App().getGreeting());

        System.out.println("args: " + Arrays.toString(args));

        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");

        Producer<String, String> producer = new KafkaProducer<>(props, new StringSerializer(), new StringSerializer());

        try {
            for (int i = 0; i < 5; i++)
                producer.send(new ProducerRecord<>("test-topic", Integer.toString(i), Integer.toString(i)));
        } catch (ProducerFencedException | OutOfOrderSequenceException | AuthorizationException e) {
            producer.close();
        }
        producer.close();

    }
}

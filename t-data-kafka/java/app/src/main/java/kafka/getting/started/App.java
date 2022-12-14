/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package kafka.getting.started;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Arrays;

public class App {
    public String getGreeting() {
        return "Hello World!";
    }

    public static void main(String[] args) {
        System.out.println(new App().getGreeting());

        System.out.println("args: " + Arrays.toString(args));

        if (args.length == 0){
            System.out.println("Invalid argument: please pass <consumer|producer>");
            System.exit(0);
        }

        String appType = args[0];
        
        if(appType.equals("producer")){
            BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
            MyProducer producer = new MyProducer("localhost:9092");
            try {
                while(true){
                    System.out.println("enter msg");
                    String msg = br.readLine();
                    System.out.println(msg);
    
                    producer.send(msg);
                }
              
            } catch (Exception e) {
                System.out.println("error while taking std input: "+e);
            }
        }else if(appType.equals("consumer")){
            MyConsumer.start("localhost:9092");
        }else{
            System.out.println("Invalid argument: please pass <consumer|producer>"+ appType);
            System.exit(0);
        }





    }
}

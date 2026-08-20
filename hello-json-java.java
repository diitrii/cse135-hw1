import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.databind.ObjectMapper;

public class HelloJSON {
    public static void main(String[] args) throws Exception {
        System.out.println("Cache-Control: no-cache\n");
        System.out.println("Content-type: application/json\n\n");

        String date = java.time.LocalDateTime.now().toString();
        String address = System.getenv("REMOTE_ADDR");

        Map<String, String> message = new HashMap<>();
        message.put("title", "Hello, Java!");
        message.put("heading", "Hello, Java!");
        message.put("message", "This page was generated with the Java programming language");
        message.put("time", date);
        message.put("IP", address);

        String json = new ObjectMapper().writeValueAsString(message);
        System.out.println(json);
    }
}


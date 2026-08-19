import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Map;

public class environmentJava {
    public static void main(String[] args) {
        System.out.println("Cache-Control: no-cache\n");
        System.out.println("Content-type: text/html \n\n");
        System.out.println("<!DOCTYPE html>"+
        "<html><head><title>Environment Variables</title>"
        "</head><body><h1 align=\"center\">Environment Variables</h1>"
        "<hr>");
        
        Map<String, String> env = System.getenv();
        List<String> variables = new ArrayList<>(env.keySet());
        Collections.sort(variables);
        for (String variable : variables) {
            System.out.println(
                "<b>" + variable + ":</b> " + env.get(variable) + "<br />"
            );
        }
        System.out.println("</body></html>");
    }
}







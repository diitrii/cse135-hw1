public class HelloHTML{
    public static void main(String[] args) throws Exception {
        System.out.println("Cache-Control: no-cache\n");
        System.out.println("Content-Type: text/html\n\n");
        System.out.println("<!DOCTYPE html>");
        System.out.println("<html>");
        System.out.println("<head>");
        System.out.println("<title>Hello CGI World</title>");
        System.out.println("</head>");
        System.out.println("<body>");
        System.out.println("<h1 align=center>Hello HTML World</h1><hr/>");
        System.out.println("<p>Hello World</p>");
        System.out.println("<p>This page was generated with the Java programming langauge</p>");
        String date = java.time.LocalDateTime.now().toString();
        System.out.println("<p>This program was generated at: " + date + "</p>");
        String address = System.getenv("REMOTE_ADDR");
        System.out.println("<p>Your current IP Address is: " + address + "</p>");
        System.out.println("</body>");
        System.out.println("</html>");
    }
}
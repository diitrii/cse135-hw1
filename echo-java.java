public class EchoJava {
    public static void main(String[] args) throws Exception {
        System.out.println("Cache-Control: no-cache\n");
        System.out.println("Content-type: text/html \n\n");
        System.out.println("<!DOCTYPE html>"+
        "<html><head><title>General Request Echo</title>"
        "</head><body><h1 align=\"center\">General Request Echo</h1>"
        "<hr>");

        System.out.println("<p><b>HTTP Protocol:</b>" + System.getenv("SERVER_PROTOCOL") + "</p>");
        System.out.println("<p><b>HTTP Method:</b>" + System.getenv("REQUEST_METHOD") + "</p>");
        System.out.println("<p><b>Query String:</b>" + System.getenv("QUERY_STRING") + "</p>");

        int contentLength = Integer.parseInt(System.getenv("CONTENT_LENGTH"));
        byte[] formData = System.in.readNBytes(contentLength);
        int bytesRead = formData.length;

        System.out.println("<p><b>Message Body:</b> " + new String(formData, 0, bytesRead) + "</p>");

        System.out.println("</body></html>");
    }
}

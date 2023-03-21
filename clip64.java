import java.util.Base64;
import java.util.Scanner;
import java.awt.datatransfer.StringSelection;
import java.awt.Toolkit;

public class clip64 {
  public static void main(String[] args){
      Scanner sc = new Scanner(System.in);
      Base64.Decoder decoder = Base64.getDecoder();

      String decodedString = new String(decoder.decode(args[0]));

      Toolkit.getDefaultToolkit()
        .getSystemClipboard()
        .setContents(
                new StringSelection(decodedString),
                null
      );

      System.out.println("[" + decodedString + "]" + " copied to clipboard");
  }
}
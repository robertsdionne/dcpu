package dcpu;

import dcpu.generated.DcpuLexer;
import dcpu.generated.DcpuParser;

import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;

import java.io.IOException;

public class Dcpu {

  public static void main(String... arguments) throws IOException {
    final ANTLRInputStream input = new ANTLRInputStream(System.in);
    final DcpuLexer lexer = new DcpuLexer(input);
    final CommonTokenStream tokens = new CommonTokenStream(lexer);
    final DcpuParser parser = new DcpuParser(tokens);
    final ParseTree tree = parser.program();
    System.out.println(tree.toStringTree(parser));
  }
}

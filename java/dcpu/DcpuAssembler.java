package dcpu;

import dcpu.generated.DcpuBaseVisitor;
import dcpu.generated.DcpuParser;
import dcpu.generated.DcpuProgram.Operand;

import org.antlr.v4.runtime.RuleContext;
import org.antlr.v4.runtime.tree.ParseTreeProperty;
import org.antlr.v4.runtime.tree.TerminalNode;

public class DcpuAssembler extends DcpuBaseVisitor<Void> {
  
  final ParseTreeProperty<Byte[]> bytes = new ParseTreeProperty<Byte[]>();
  final ParseTreeProperty<Boolean> pending = new ParseTreeProperty<Boolean>();
  final ParseTreeProperty<Integer> argumentValue = new ParseTreeProperty<Integer>();
  final ParseTreeProperty<Boolean> argumentHasPayload = new ParseTreeProperty<Boolean>();
  final ParseTreeProperty<Boolean> argumentPayloadNeedsResolution = new ParseTreeProperty<Boolean>();
  final ParseTreeProperty<String> argumentPayloadLabel = new ParseTreeProperty<String>();
  final ParseTreeProperty<Integer> argumentPayload = new ParseTreeProperty<Integer>();

  @Override public Void visitArgumentA(DcpuParser.ArgumentAContext context) {
    if (context.REGISTER() != null) {
      int value = registerToValue(context.REGISTER());
      argumentValue.put(context, value);
    } else if (context.POP() != null) {
      argumentValue.put(context, Operand.Type.PUSH_POP_VALUE);
    } else if (context.PEEK() != null) {
      argumentValue.put(context, Operand.Type.PEEK_VALUE);
    } else if (context.STACK_POINTER() != null) {
      argumentValue.put(context, Operand.Type.STACK_POINTER_VALUE);
    } else if (context.PROGRAM_COUNTER() != null) {
      argumentValue.put(context, Operand.Type.PROGRAM_COUNTER_VALUE);
    } else if (context.EXTRA() != null) {
      argumentValue.put(context, Operand.Type.EXTRA_VALUE);
    }
    this.visitChildren(context);
    System.out.print(argumentValue.get(context));
    System.out.print(" ");
    if (argumentHasPayload.get(context) != null && argumentHasPayload.get(context)) {
      if (argumentPayloadNeedsResolution.get(context) != null &&
          argumentPayloadNeedsResolution.get(context)) {
        System.out.println(argumentPayloadLabel.get(context));
      } else {
        System.out.println(argumentPayload.get(context));
      }
    }
    System.out.println();
    return null;
  }
  
  @Override public Void visitArgumentB(DcpuParser.ArgumentBContext context) {
    if (context.REGISTER() != null) {
      int value = registerToValue(context.REGISTER());
      argumentValue.put(context, value);
    } else if (context.PUSH() != null) {
      argumentValue.put(context, Operand.Type.PUSH_POP_VALUE);
    } else if (context.PEEK() != null) {
      argumentValue.put(context, Operand.Type.PEEK_VALUE);
    } else if (context.STACK_POINTER() != null) {
      argumentValue.put(context, Operand.Type.STACK_POINTER_VALUE);
    } else if (context.PROGRAM_COUNTER() != null) {
      argumentValue.put(context, Operand.Type.PROGRAM_COUNTER_VALUE);
    } else if (context.EXTRA() != null) {
      argumentValue.put(context, Operand.Type.EXTRA_VALUE);
    }
    this.visitChildren(context);
    System.out.print(argumentValue.get(context));
    System.out.print(" ");
    if (argumentHasPayload.get(context) != null && argumentHasPayload.get(context)) {
      if (argumentPayloadNeedsResolution.get(context) != null &&
          argumentPayloadNeedsResolution.get(context)) {
        System.out.println(argumentPayloadLabel.get(context));
      } else {
        System.out.println(argumentPayload.get(context));
      }
    }
    System.out.println();
    return null;
  }

  private int registerToValue(TerminalNode register) {
    return Operand.Register.valueOf(register.getText().toUpperCase()).getNumber();
  }
  
  @Override public Void visitLocationInRegister(DcpuParser.LocationInRegisterContext context) {
    argumentValue.put(context.getParent(),
        Operand.Type.LOCATION_IN_REGISTER_VALUE + registerToValue(context.REGISTER()));
    return null;
  }
  
  @Override public Void visitLocationOffsetByRegister(DcpuParser.LocationOffsetByRegisterContext context) {
    final RuleContext parent = context.getParent();
    argumentValue.put(parent,
        Operand.Type.LOCATION_OFFSET_BY_REGISTER_VALUE + registerToValue(context.REGISTER()));
    argumentHasPayload.put(parent, true);
    if (context.value().IDENTIFIER() != null) {
      argumentPayloadNeedsResolution.put(parent, true);
      argumentPayloadLabel.put(parent, context.value().IDENTIFIER().getText());
    } else if (context.value().NUMBER() != null) {
      argumentPayload.put(parent, parseNumber(context.value().NUMBER()));
    }
    return null;
  }
  
  @Override public Void visitPick(DcpuParser.PickContext context) {
    final RuleContext parent = context.getParent();
    argumentValue.put(parent, Operand.Type.PICK_VALUE);
    argumentHasPayload.put(parent, true);
    argumentPayload.put(parent, parseNumber(context.NUMBER()));
    return null;
  }
  
  @Override public Void visitLocation(DcpuParser.LocationContext context) {
    final RuleContext parent = context.getParent();
    argumentValue.put(parent, Operand.Type.LOCATION_VALUE);
    argumentHasPayload.put(parent, true);
    if (context.value().IDENTIFIER() != null) {
      argumentPayloadNeedsResolution.put(parent, true);
      argumentPayloadLabel.put(parent, context.value().IDENTIFIER().getText());
    } else if (context.value().NUMBER() != null) {
      argumentPayload.put(parent, parseNumber(context.value().NUMBER()));
    }
    return null;
  }
  
  @Override public Void visitValue(DcpuParser.ValueContext context) {
    final RuleContext parent = context.getParent();
    if (parent.getRuleIndex() == DcpuParser.RULE_argumentA ||
        parent.getRuleIndex() == DcpuParser.RULE_argumentB) {
      if (context.IDENTIFIER() != null) {
        argumentValue.put(parent, Operand.Type.LITERAL_VALUE);
        argumentHasPayload.put(parent, true);
        argumentPayloadNeedsResolution.put(parent, true);
        argumentPayloadLabel.put(parent, context.IDENTIFIER().getText());
      } else if (context.NUMBER() != null) {
        int value = parseNumber(context.NUMBER());
        if (value == 0xffff || value == -1) {
          argumentValue.put(parent, 0x20);
        } else if (value < 31) {
          argumentValue.put(parent, 0x21 + value);
        } else {
          argumentValue.put(parent, Operand.Type.LITERAL_VALUE);
          argumentHasPayload.put(parent, true);
          argumentPayload.put(parent, value);
        }
      }
    }
    return null;
  }

  private Integer parseNumber(TerminalNode number) {
    if (number.getText().startsWith("0x")) {
      return Integer.valueOf(number.getText().substring(2), 16);
    } else {
      return Integer.valueOf(number.getText());
    }
  }
}

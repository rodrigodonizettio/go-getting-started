package br.com.rodrigodonizettio.pcbook.serializer;

import br.com.rodrigodonizettio.pcbook.pb.Laptop;
import com.google.protobuf.util.JsonFormat;

import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class Serializer {
  public static void main(String[] args) throws IOException {
    Serializer serializer = new Serializer();
    Laptop laptop = serializer.readBinaryFile("laptop.bin");
    serializer.writeJsonFile(laptop, "laptopJson.json");
  }

  public void writeBinaryFile(Laptop laptop, String filename) throws IOException {
    FileOutputStream fileOutputStream = new FileOutputStream(filename);
    laptop.writeTo(fileOutputStream);
    fileOutputStream.close();
  }

  public Laptop readBinaryFile(String filename) throws IOException {
    FileInputStream fileInputStream = new FileInputStream(filename);
    Laptop laptop = Laptop.parseFrom(fileInputStream);
    fileInputStream.close();
    return laptop;
  }

  public void writeJsonFile(Laptop laptop, String filename) throws IOException {
    JsonFormat.Printer printer = JsonFormat.printer()
//          .printingEnumsAsInts()
//            .preservingProtoFieldNames()
            .includingDefaultValueFields();

    String jsonString = printer.print(laptop);

    FileOutputStream fileOutputStream = new FileOutputStream(filename);
    fileOutputStream.write(jsonString.getBytes());
    fileOutputStream.close();
  }
}

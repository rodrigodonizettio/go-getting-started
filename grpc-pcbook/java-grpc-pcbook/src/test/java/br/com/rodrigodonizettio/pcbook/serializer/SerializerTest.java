package br.com.rodrigodonizettio.pcbook.serializer;

import br.com.rodrigodonizettio.pcbook.pb.Laptop;
import br.com.rodrigodonizettio.pcbook.sample.Generator;
import org.junit.Assert;
import org.junit.Test;

import java.io.IOException;

import static org.junit.Assert.*;

public class SerializerTest {

  @Test
  public void writeAndReadBinaryFile() throws IOException {
    String binaryFile = "laptop.bin";
    Laptop laptop1 = new Generator().newLaptop();

    Serializer serializer = new Serializer();
    serializer.writeBinaryFile(laptop1, binaryFile);

    Laptop laptop2 = serializer.readBinaryFile(binaryFile);
    Assert.assertEquals(laptop1, laptop2);
  }
}
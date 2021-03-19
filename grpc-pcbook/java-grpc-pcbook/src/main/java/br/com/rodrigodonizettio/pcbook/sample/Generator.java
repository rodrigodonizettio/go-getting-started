package br.com.rodrigodonizettio.pcbook.sample;

import br.com.rodrigodonizettio.pcbook.pb.Cpu;
import br.com.rodrigodonizettio.pcbook.pb.Gpu;
import br.com.rodrigodonizettio.pcbook.pb.Keyboard;
import br.com.rodrigodonizettio.pcbook.pb.Laptop;
import br.com.rodrigodonizettio.pcbook.pb.Memory;
import br.com.rodrigodonizettio.pcbook.pb.Screen;
import br.com.rodrigodonizettio.pcbook.pb.Storage;
import com.google.protobuf.Timestamp;

import java.time.Instant;
import java.util.Random;

public class Generator {
  private Random random;

  public Generator() {
    this.random = new Random();
  }

  public static void main(String[] args) {
    Generator generator = new Generator();
    Laptop laptop = generator.newLaptop();
    System.out.println(laptop);
  }

  public Keyboard newKeyboard() {
    return Keyboard.newBuilder()
            .setLayout(randomKeyboardLayout())
            .setBacklit(random.nextBoolean())
            .build();
  }

  public Cpu newCpu() {
    String brand = randomCpuBrand();
    String name = randomCpuName(brand);
    int numberCores = randomInt(2, 8);
    int numberThreads = randomInt(numberCores, 12);
    double minGhz = randomDouble(1.0, 4.5);
    double maxGhz = randomDouble(minGhz, 5.5);

    return Cpu.newBuilder()
            .setBrand(brand)
            .setName(name)
            .setNumberCores(numberCores)
            .setNumberThreads(numberThreads)
            .setMinGhz(minGhz)
            .setMaxGhz(maxGhz)
            .build();
  }

  public Gpu newGpu() {
    String brand = randomGpuBrand();
    String name = randomGpuName(brand);
    double minGhz = randomDouble(2.0, 5.5);
    double maxGhz = randomDouble(minGhz, 6.5);

    Memory memory = Memory.newBuilder()
            .setValue(randomInt(1, 6))
            .setUnit(Memory.Unit.GIGABYTE)
            .build();

    return Gpu.newBuilder()
            .setBrand(brand)
            .setName(name)
            .setMinGhz(minGhz)
            .setMaxGhz(maxGhz)
            .setMemory(memory)
            .build();
  }

  public Memory newRam() {
    return Memory.newBuilder()
            .setValue(randomInt(8, 32))
            .setUnit(Memory.Unit.GIGABYTE)
            .build();
  }

  public Storage newSsd() {
    Memory memory = Memory.newBuilder()
            .setValue(randomInt(128, 2048))
            .setUnit(Memory.Unit.GIGABYTE)
            .build();

    return Storage.newBuilder()
            .setDriver(Storage.Driver.SSD)
            .setMemory(memory)
            .build();
  }

  public Storage newHdd() {
    Memory memory = Memory.newBuilder()
            .setValue(randomInt(1, 5))
            .setUnit(Memory.Unit.TERABYTE)
            .build();

    return Storage.newBuilder()
            .setDriver(Storage.Driver.HDD)
            .setMemory(memory)
            .build();
  }

  public Screen newScreen() {
    int height = randomInt(800, 5000);
    int width = height * (16 / 9);

    Screen.Resolution resolution = Screen.Resolution.newBuilder()
            .setHeight(height)
            .setWidth(width)
            .build();

    return Screen.newBuilder()
            .setSizeInch(randomFloat(14, 28))
            .setResolution(resolution)
            .setPanel(randomScreenPanel())
            .setMultitouch(random.nextBoolean())
            .build();
  }

  public Laptop newLaptop() {
    String brand = randomLaptopBrand();
    String name = randomLaptopName(brand);
    double weightKg = randomDouble(0.5, 2.5);
    double priceBrl = randomDouble(4000, 12000);
    int releaseYear = randomInt(2015, 2021);

    return Laptop.newBuilder()
            .setBrand(brand)
            .setName(name)
            .setWeightKg(weightKg)
            .setPriceBrl(priceBrl)
            .setReleaseYear(releaseYear)
            .setCpu(newCpu())
            .setRam(newRam())
            .setScreen(newScreen())
            .setKeyboard(newKeyboard())
            .addGpus(newGpu())
            .addStorages(newSsd())
            .addStorages(newHdd())
            .setUpdatedAt(timestampNow())
            .build();
  }

  private Timestamp timestampNow() {
    Instant now = Instant.now();
    return Timestamp.newBuilder()
            .setSeconds(now.getEpochSecond())
            .setNanos(now.getNano())
            .build();
  }

  private String randomLaptopName(String brand) {
    switch(brand) {
      case "Asus":
        return randomStringFromSet("ROG", "VivoBook", "ZenBook");
      case "Dell":
        return randomStringFromSet("Alienware", "Inspiron", "Vostro");
      default:
        return randomStringFromSet("IdeaPad", "Legion", "ThinkPad");
    }
  }

  private String randomLaptopBrand() {
    return randomStringFromSet("Asus", "Dell", "Lenovo");
  }

  private Screen.Panel randomScreenPanel() {
    if(random.nextBoolean()) {
      return Screen.Panel.IPS;
    }
    return Screen.Panel.OLED;
  }

  private float randomFloat(float min, float max) {
    return min + random.nextFloat() * (max - min);
  }

  private String randomGpuName(String brand) {
    if(brand.equals("NVidia")) {
      return randomStringFromSet("930", "1080", "2060");
    }
    return randomStringFromSet("580, 590, Vega");
  }

  private String randomGpuBrand() {
    return randomStringFromSet("NVidia", "AMD");
  }

  private double randomDouble(double min, double max) {
    return min + random.nextDouble() * (max - min);
  }

  private int randomInt(int min, int max) {
    return min + random.nextInt(max - min + 1);
  }

  private String randomCpuName(String brand) {
    if(brand.equals("Intel")) {
      return randomStringFromSet("Xeon", "Core i3", "Core i5", "Core i7", "Core i9");
    }
    return randomStringFromSet("Ryzen 3", "Ryzen 5", "Ryzen 7");
  }

  private String randomCpuBrand() {
    return randomStringFromSet("Intel", "AMD");
  }

  private String randomStringFromSet(String ... s) {
    int length = s.length;
    if(length == 0) {
      return "";
    }
    return s[random.nextInt(length)];
  }

  private Keyboard.Layout randomKeyboardLayout() {
    switch(random.nextInt(3)) {
      case 1:
        return Keyboard.Layout.QWERTY;
      case 2:
        return Keyboard.Layout.QWERTZ;
      default:
        return Keyboard.Layout.AZERTY;
    }
  }
}

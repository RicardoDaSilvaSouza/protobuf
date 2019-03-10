package br.com.example.protobuf

import example.simple.Simple.SimpleMessage
import java.io.File
import java.io.FileInputStream
import java.io.FileOutputStream

fun main(args: Array<String>) {
    println("Hello World")

    val builder: SimpleMessage.Builder = SimpleMessage.newBuilder()

    builder.id = 42
    builder.isSimple = true
    builder.name = "Ricardo"

    builder.addSampleList(1)
    builder.addSampleList(2)
    builder.addSampleList(3)

    builder.addAllSampleList(MutableList(10) {it + 1})

    var message = builder.build()

    println("""Writing to file...
            |${message.toString()}""")

    var file  = File("/Users/ricardo.souza/_file_test.bin")

    file.delete()

    file.createNewFile().also { if(it) FileOutputStream(file)
            .use { message.writeTo(it) }}

    SimpleMessage.parseFrom(FileInputStream(file)).let { println("""Read from file...
        |$it""".trimMargin()) }

}
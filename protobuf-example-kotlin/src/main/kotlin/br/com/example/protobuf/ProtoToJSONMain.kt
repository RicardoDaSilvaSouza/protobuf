package br.com.example.protobuf

import com.google.protobuf.util.JsonFormat
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

//    var message = builder.build()
// Protobuf to json
    var json = JsonFormat.printer().print(builder)
    println("Message as JSON: $json")

//  Json to Protobuf

    val newBuilder = SimpleMessage.newBuilder()

    JsonFormat.parser().ignoringUnknownFields().merge(json, newBuilder)
    println("Message as Protobuf: $newBuilder")

}
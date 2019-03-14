package br.com.example.protobuf

import example.complex.Complex.*
import java.util.*

fun main(args: Array<String>) {
    println("Complex example")

    var dummyMessage = createDummyMessage(1, "Dummy message 1")

    var complexMessage = ComplexMessage.newBuilder()
    complexMessage.setDummyMessage(dummyMessage)

    complexMessage.addAllMultipleDumy(Arrays.asList(
        createDummyMessage(1, "Dummy message 1"),
        createDummyMessage(2, "Dummy message 2"),
        createDummyMessage(3, "Dummy message 3")
    ))

    var message = complexMessage.build()

    println(message.toString())

}

fun createDummyMessage(id: Int, name: String): DummyMessage {
    return DummyMessage.newBuilder()
            .setId(id)
            .setName(name)
            .build()
}
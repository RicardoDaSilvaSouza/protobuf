package br.com.example.protobuf

import example.enumeration.EnumExample

fun main(args: Array<String>) {
    println("Example for Enums")

    val builder = EnumExample.EnumMessage.newBuilder()

    builder.id = 1
    builder.dayOfTheWeek = EnumExample.DayOfTheWeek.WEDNESDAY

    var message = builder.build()

    println(message)
}
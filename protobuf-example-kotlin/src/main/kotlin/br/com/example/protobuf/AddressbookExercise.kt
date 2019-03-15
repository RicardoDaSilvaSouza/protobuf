package br.com.example.protobuf

import com.example.tutorial.AddressBookProtos.*
import java.io.FileInputStream
import java.io.FileOutputStream
import java.util.*
import kotlin.collections.ArrayList

fun main(args: Array<String>) {
    val addressBookBuilder = AddressBook.newBuilder()

    addressBookBuilder.mergeFrom(FileInputStream(args[0]))

    addressBookBuilder.addAllPeople(generatePeople(3))

    val addressBook = addressBookBuilder.build()

    addressBook.writeTo(FileOutputStream(args[0]))

    println(addressBook)
}

private fun generatePeople(qty: Int): List<Person> {
    val result = ArrayList<Person>()
    val personBuilder = Person.newBuilder()
    val phoneNumber = Person.PhoneNumber.newBuilder()
    val timeStampBuilder = personBuilder.lastUpdatedBuilder

    for(i in 0..qty){

        timeStampBuilder.seconds = Calendar.getInstance().timeInMillis
        personBuilder.lastUpdated = timeStampBuilder.build()

        personBuilder.id = 1
        personBuilder.name = "User$i name"
        personBuilder.email = "user$i.email@mail.com"

        phoneNumber.number = "119853475$i"
        phoneNumber.type = Person.PhoneType.MOBILE

        val phoneNumberMob = phoneNumber.build()

        personBuilder.clearPhones()
        personBuilder.addPhones(phoneNumberMob)

        result.add(personBuilder.build())
    }

    return result
}
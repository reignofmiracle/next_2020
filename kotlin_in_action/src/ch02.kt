package ch02

fun main(args: Array<String>) {
//    main1(args)
    main2(args)
}

fun main1(args: Array<String>) {
    println("Hello, world!")

    println(max(10, 11))
    println(max2(10, 13))
}

fun max(a: Int, b: Int): Int {
    println("my max")
    return if (a > b) a else b
}

fun max2(a: Int, b: Int): Int = if (a > b) a else b

fun main2(args: Array<String>) {
    val name = if (args.size > 0) args[0] else "Kotlin"
    println("${name}")
}
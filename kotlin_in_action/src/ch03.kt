@file:JvmName("test")

package ch03

fun main(args: Array<String>) {
    main1(args)
    main2(args)
}

fun main1(args: Array<String>) {
    val set = hashSetOf(1, 7, 3213)
    val list = arrayListOf(1,2,3,4,5)
    val map = hashMapOf(1 to "one", 7 to "seven", 0 to "zero")

    println(set.javaClass)
    println(list.javaClass)
    println(map.javaClass)

    val strings = listOf("f", "s", "t")
    println(strings.last())

    val numbers = setOf(1,2,3,4,5)
    println(numbers.max())
}

fun main2(args: Array<String>) {
    val list = listOf(1,2,3,4,5)
    println(list.joinToString("; ", "(", ")"))
    println(list.joinToString())
    println(list.joinToString("; "))
    println(list.joinToString(prefix = "+ "))
}

fun<T> Collection<T>.joinToString(
    separator: String = ", ",
    prefix: String = "",
    postfix: String = ""
): String {
val result = StringBuilder(prefix)
    var separatorTmp = ""
    for ((index, element) in this.withIndex()) {
        result.append(separatorTmp)
        result.append(element)
        separatorTmp = separator
    }

    result.append(postfix)
    return result.toString()
}
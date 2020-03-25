package ch04

interface Clickable {
    fun click()
    fun showOff() = println("I'm clickable!")
}

interface Focusable {
    fun setFocus(b: Boolean) = println("I ${if (b) "got" else "lost"} focus")
    fun showOff() = println("I'm focusable!")
}

class Button : Clickable, Focusable {
    override fun click() {
        println("I was clicked")
    }

    override fun showOff() {
        super<Clickable>.showOff()
        super<Focusable>.showOff()
    }
}

class User(val nickname: String,
           val isSubscribed: Boolean = true)

class User2 constructor(val nickname: String,
           val isSubscribed: Boolean = true) {
    val nickname2: String
    init {
        nickname2 = nickname
    }
}

class DelegatingCollection<T>(
    innerList: Collection<T> = ArrayList<T>()
) : Collection<T> by innerList


fun main(args: Array<String>) {
    val button = Button()
    button.showOff()
    button.setFocus(true)
    button.click()

    val test1 = User("test1", false)
    val test2 = User("test2", true)
    println(test1.nickname)
    println(test2.nickname)

    val test3 = User2("test1", false)
    val test4 = User2("test2", true)
    println(test3.nickname2)
    println(test4.nickname2)
}
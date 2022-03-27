import kotlin.system.measureTimeMillis

fun collatz(n: Long) : Long {
    var count = 1L
    var m = n
    while (m > 1L) {
        if (m % 2L == 0L) {
            m /= 2L
        } else {
            m = m * 3L + 1L
        }
        count++
    }
    return count
}

fun main() {
    val time = measureTimeMillis {
        var max = 0L
        var key = 0L
        val limit = 100_000_000L
        for (i in 2L..limit) {
            var rc = collatz(i)
            if (rc > max) {
                max = rc
                key = i
            }
        }
        println("max = ${key}(${max})")
    }
    println("$time ms")

}

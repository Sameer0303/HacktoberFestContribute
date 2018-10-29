fun main(args: Array<String>) {
    var removedVowels = Regex("a|e|i|o|u").replace(args[0], "") 
    println(removedVowels)
}

function shortcut(string) {
  textSplitted = string.split("")
  result = ""
  vowel = ['a', 'i', 'u', 'e', 'o']
  
  textSplitted.forEach(item => {
    if (!vowel.includes(item)) {
      result += item    
    }
  });
  return result  
}

document.getElementById("result").innerHTML = shortcut("test")


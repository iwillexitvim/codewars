// ROT13 is a simple letter substitution cipher that replaces a letter with the letter 13 letters after it in the alphabet. ROT13 is an example of the Caesar cipher.

// Create a function that takes a string and returns the string ciphered with Rot13. If there are numbers or special characters included in the string, they should be returned as they are. Only letters from the latin/english alphabet should be shifted, like in the original Rot13 "implementation".

// test => grfg
// Test => Grfg

function rot13(message){
  const alphabet = 'abcdefghijklmnopqrstuvwxyz'.split('');
  return message.replace(/[a-zA-Z]/g, (match) => {
    const matchPosition = alphabet.findIndex((alph) => alph === match.toLowerCase());
    const nextPosition = (matchPosition + 13) % alphabet.length;
    return match.toLowerCase() === match ? alphabet[nextPosition] : alphabet[nextPosition].toUpperCase() 
  });
}

console.log(rot13("test"))
console.log(rot13("Test"))

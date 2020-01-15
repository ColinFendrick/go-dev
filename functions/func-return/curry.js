const curry = a => b => c => d => `${a} ${b} ${c} ${d}`;

console.log(
  curry("curry")("functions")("are")("delicious")
)
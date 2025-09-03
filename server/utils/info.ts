import * as cfonts from "cfonts";
const colorOptions = [
   'yellow',
  'cyan', 'white', 'gray'
];
function getRandomColors() {
  const shuffled = [...colorOptions].sort(() => 0.5 - Math.random());
  return shuffled.slice(0, Math.floor(Math.random() * 2) + 1);
}
(cfonts as any).default.say('KeyMesh', {
	font: 'block',              // define the font face
	align: 'left',              // define text alignment
  colors: getRandomColors(),    // 随机颜色
	background: 'transparent',  // define the background color, you can also use `backgroundColor` here as key
	letterSpacing: 1,           // define letter spacing
	lineHeight: 1,              // define the line height
	space: true,                // define if the output text should have empty lines on top and on the bottom
	maxLength: '0',             // define how many character can be on one line
	gradient: false,            // define your two gradient colors
	independentGradient: false, // define if you want to recalculate the gradient for each new line
	transitionGradient: false,  // define if this is a transition between colors directly
	rawMode: false,             // define if the line breaks should be CRLF (`\r\n`) over the default LF (`\n`)
	env: 'node'                 // define the environment cfonts is being executed in
});
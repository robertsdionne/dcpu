

var canvas, gl,
    dcpu, display,
    drawable, beamDrawable, debugDrawable,
    buffer, beamBuffer, debugBuffer,
    program,
    projection, modelView;


var load = function() {
  canvas = document.getElementById('c');
  gl = canvas.getContext('experimental-webgl');
  setup();
  requestAnimationFrame(animate);
};
window.addEventListener('load', load);


var animate = function() {
  update();
  draw();
  requestAnimationFrame(animate);
};

var x = 0;
var y = 0;
var z = 0;
var vx = 1;
var vy = 2;
var vz = 3;


var randomUint16Array = function() {
  x += vx;
  y += vy;
  z += vz;
  if (x < 0) {
    x = 0;
    vx *= -1;
  }
  if (x > 223) {
    x = 223;
    vx *= -1;
  }if (y < 0) {
    y = 0;
    vy *= -1;
  }
  if (y > 223) {
    y = 223;
    vy *= -1;
  }
  if (z < 0) {
    z = 0;
    vz *= -1;
  }
  if (z > 223) {
    z = 223;
    vz *= -1;
  }
  var word0 = (y << 8) | x;
  var word1 = z;
  return new Uint16Array([
    0x0000 + word0, 0x0700 + word1,
    0x2000 + word0, 0x0700 + word1,
    0x2020 + word0, 0x0700 + word1,
    0x0020 + word0, 0x0700 + word1,
    0x0000 + word0, 0x0700 + word1,

    0x0000 + word0, 0x0720 + word1,
    0x0000 + word0, 0x0700 + word1,
    0x2000 + word0, 0x0700 + word1,
    0x2000 + word0, 0x0720 + word1,
    0x2020 + word0, 0x0720 + word1,
    0x2020 + word0, 0x0700 + word1,
    0x0020 + word0, 0x0700 + word1,
    0x0020 + word0, 0x0720 + word1,

    0x0000 + word0, 0x0720 + word1,
    0x2000 + word0, 0x0720 + word1,
    0x2020 + word0, 0x0720 + word1,
    0x0020 + word0, 0x0720 + word1,
    0x0000 + word0, 0x0720 + word1,
    0x0000 + word0, 0x0020 + word1,

    0x3000 + word0, 0x0000 + word1,
    0x3000 + word0, 0x0700 + word1,
    0x5000 + word0, 0x0700 + word1,
    0x5020 + word0, 0x0700 + word1,
    0x3020 + word0, 0x0700 + word1,
    0x3000 + word0, 0x0700 + word1,

    0x3000 + word0, 0x0720 + word1,
    0x3000 + word0, 0x0700 + word1,
    0x5000 + word0, 0x0700 + word1,
    0x5000 + word0, 0x0720 + word1,
    0x5020 + word0, 0x0720 + word1,
    0x5020 + word0, 0x0700 + word1,
    0x3020 + word0, 0x0700 + word1,
    0x3020 + word0, 0x0720 + word1,

    0x3000 + word0, 0x0720 + word1,
    0x5000 + word0, 0x0720 + word1,
    0x5020 + word0, 0x0720 + word1,
    0x3020 + word0, 0x0720 + word1,
    0x3000 + word0, 0x0720 + word1,
    0x3000 + word0, 0x0020 + word1,

    0x0000, 0x0000,
    0x0000, 0x0700,
    0x2000, 0x0700,
    0x2000, 0x0000,
    0xDF00, 0x0000,
    0xDF00, 0x0700,
    0xFF00, 0x0700,
    0xFF20, 0x0700,
    0xFF20, 0x0000,
    0xFFDF, 0x0000,
    0xFFDF, 0x0700,
    0xFFFF, 0x0700,
    0xDFFF, 0x0700,
    0xDFFF, 0x0000,
    0x20FF, 0x0000,
    0x20FF, 0x0700,
    0x00FF, 0x0700,
    0x00DF, 0x0700,
    0x00DF, 0x0000,
    0x0020, 0x0000,
    0x0020, 0x0700,
    0x0000, 0x0700,
    0x0000, 0x0000,

    0x0000, 0x00FF,
    0x0000, 0x07FF,
    0x2000, 0x07FF,
    0x2000, 0x00FF,
    0xDF00, 0x00FF,
    0xDF00, 0x07FF,
    0xFF00, 0x07FF,
    0xFF20, 0x07FF,
    0xFF20, 0x00FF,
    0xFFDF, 0x00FF,
    0xFFDF, 0x07FF,
    0xFFFF, 0x07FF,
    0xDFFF, 0x07FF,
    0xDFFF, 0x00FF,
    0x20FF, 0x00FF,
    0x20FF, 0x07FF,
    0x00FF, 0x07FF,
    0x00DF, 0x07FF,
    0x00DF, 0x00FF,
    0x0020, 0x00FF,
    0x0020, 0x07FF,
    0x0000, 0x07FF
  ]);
};


var getFrustumMatrix = function(
    left, right, bottom, top, near, far) {
  var a = (right + left) / (right - left);
  var b = (top + bottom) / (top - bottom);
  var c = -(far + near) / (far - near);
  var d = -(2 * far * near) / (far - near);
  return [
    2 * near / (right - left), 0, 0, 0,
    0, 2 * near / (top - bottom), 0, 0,
    a, b, c, -1,
    0, 0, d, 0
  ];
};


var setup = function() {
  gl.viewport(0, 0, canvas.width, canvas.height);
  gl.clearColor(0.0, 0.0, 0.0, 1.0);
  gl.enable(gl.BLEND);
  gl.depthMask(gl.FALSE);
  gl.disable(gl.CULL_FACE);
  gl.blendFunc(gl.SRC_ALPHA, gl.ONE);
  gl.lineWidth(2.0);

  var aspect = canvas.width / canvas.height;
  projection = getFrustumMatrix(-0.1 * aspect, 0.1 * aspect, -0.1, 0.1, 0.1, 1000);
  modelView = [
    1.0, 0.0, 0.0, 0.0,
    0.0, 1.0, 0.0, 0.0,
    0.0, 0.0, 1.0, 0.0,
    0.0, -0.8, -1.5, 1.0
  ];

  var vertex = gl.createShader(gl.VERTEX_SHADER);
  var vertex_source = document.getElementById('v0').text;
  gl.shaderSource(vertex, vertex_source);
  gl.compileShader(vertex);
  if (!gl.getShaderParameter(vertex, gl.COMPILE_STATUS)) {
    var log = gl.getShaderInfoLog(vertex);
    gl.deleteShader(vertex);
    vertex = null;
    throw new Error('v0: ' + log + vertex_source);
  }

  var fragment = gl.createShader(gl.FRAGMENT_SHADER);
  var fragment_source = document.getElementById('f0').text;
  gl.shaderSource(fragment, fragment_source);
  gl.compileShader(fragment);
  if (!gl.getShaderParameter(fragment, gl.COMPILE_STATUS)) {
    var log = gl.getShaderInfoLog(fragment);
    gl.deleteShader(fragment);
    fragment = null;
    throw new Error('f0: ' + log + fragment_source);
  }

  program = gl.createProgram();
  gl.attachShader(program, vertex);
  gl.attachShader(program, fragment);
  gl.linkProgram(program);
  if (!gl.getProgramParameter(program, gl.LINK_STATUS)) {
    var log = gl.getProgramInfoLog(program);
    gl.detachShader(program, vertex);
    gl.deleteShader(vertex);
    vertex = null;
    gl.detachShader(program, fragment);
    gl.deleteShader(fragment);
    fragment = null;
    gl.deleteProgram(program);
    program = null;
    throw new Error(log);
  }

  dcpu = {
    memory: randomUint16Array()
  };
  display = new Display(dcpu);
  display.map(0, dcpu.memory.length / 2);

  drawable = display.lineView(gl);
  buffer = gl.createBuffer();
  gl.bindBuffer(gl.ARRAY_BUFFER, buffer);
  gl.bufferData(gl.ARRAY_BUFFER, drawable.data, gl.STREAM_DRAW);

  beamDrawable = display.view(gl);
  beamBuffer = gl.createBuffer();
  gl.bindBuffer(gl.ARRAY_BUFFER, beamBuffer);
  gl.bufferData(gl.ARRAY_BUFFER, beamDrawable.data, gl.STREAM_DRAW);

  debugDrawable = display.debugView(gl);
  debugBuffer = gl.createBuffer();
  gl.bindBuffer(gl.ARRAY_BUFFER, debugBuffer);
  gl.bufferData(gl.ARRAY_BUFFER, debugDrawable.data, gl.STREAM_DRAW);
};


var update = function() {
  dcpu.memory = randomUint16Array();
  display.update();
  console.log(display.currentState);

  drawable = display.lineView(gl);
  gl.bindBuffer(gl.ARRAY_BUFFER, buffer);
  gl.bufferData(gl.ARRAY_BUFFER, drawable.data, gl.STREAM_DRAW);

  beamDrawable = display.view(gl);
  gl.bindBuffer(gl.ARRAY_BUFFER, beamBuffer);
  gl.bufferData(gl.ARRAY_BUFFER, beamDrawable.data, gl.STREAM_DRAW);

  debugDrawable = display.debugView(gl);
  gl.bindBuffer(gl.ARRAY_BUFFER, debugBuffer);
  gl.bufferData(gl.ARRAY_BUFFER, debugDrawable.data, gl.STREAM_DRAW);
};


var drawBuffer = function(drawable, buffer) {
  gl.bindBuffer(gl.ARRAY_BUFFER, buffer);
  var position = gl.getAttribLocation(program, 'position');
  var color = gl.getAttribLocation(program, 'color');
  gl.enableVertexAttribArray(position);
  gl.vertexAttribPointer(position, 3, gl.FLOAT, false, 28, 0);
  gl.enableVertexAttribArray(color);
  gl.vertexAttribPointer(color, 4, gl.FLOAT, false, 28, 12);
  gl.drawArrays(drawable.elementType, 0, drawable.elementCount);
  gl.disableVertexAttribArray(position);
  gl.disableVertexAttribArray(color);
};


var draw = function() {
  gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

  gl.useProgram(program);
  gl.uniformMatrix4fv(gl.getUniformLocation(program, 'projection'), false, projection);
  gl.uniformMatrix4fv(gl.getUniformLocation(program, 'model_view'), false, modelView);

  drawBuffer(beamDrawable, beamBuffer);
  drawBuffer(drawable, buffer);
  drawBuffer(debugDrawable, debugBuffer);

  gl.flush();
}

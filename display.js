

var Display = function(dcpu) {
  this.dcpu = dcpu;
  this.angle = 0;
  this.currentState = Display.State.NO_DATA;
  this.lastError = Display.Error.NONE;
  this.memoryMapOffset = 0x0000;
  this.numberOfVertices = 0;
  this.targetAngle = 90;
};


Display.ANGULAR_VELOCITY = 50; // degrees per second


Display.BYTE_MAXIMUM = 255;


Display.COORDINATE_MASK = 0x00FF;


Display.COLOR_MASK = 0x0007;


Display.COLOR_SHIFT = 8;


Display.INTENSITY_MASK = 0x0001;


Display.INTENSITY_SHIFT = 9;


Display.LERP_ALPHA = 0.01;


Display.TURN_EPSILON = 1e-2;


Display.Y_SHIFT = 8;


Display.Error = {
  NONE: 0x0000,
  BROKEN: 0xFFFF
};


Display.State = {
  NO_DATA: 0x0000,
  RUNNING: 0x0001,
  TURNING: 0x0002
};


Display.prototype.buildVertex = function(word0, word1, opt_alpha) {
  var x = (word0 & Display.COORDINATE_MASK) / Display.BYTE_MAXIMUM - 0.5;
  var z = (word0 >> Display.Y_SHIFT) / Display.BYTE_MAXIMUM - 0.5;
  var y = (word1 & Display.COORDINATE_MASK) / Display.BYTE_MAXIMUM + 0.5;
  var color = (word1 >> Display.COLOR_SHIFT) & Display.COLOR_MASK;
  var r = (0x4 & color) ? 1.0 : 0.0;
  var g = (0x2 & color) ? 1.0 : 0.0;
  var b = (0x1 & color) ? 1.0 : 0.0;
  var a = opt_alpha || 0.8;
  var theta = Math.PI / 180.0 * this.angle;
  var cos0 = Math.cos(theta);
  var sin0 = Math.sin(theta);
  return [
    x * cos0 + z * sin0, y, x * sin0 - z * cos0, r, g, b, a
  ];
};


Display.prototype.lineView = function(gl) {
  var data = [];
  for (var i = 0; i < this.numberOfVertices; ++i) {
    Array.prototype.push.apply(
        data, this.buildVertex(this.dcpu.memory[2 * i], this.dcpu.memory[2 * i + 1]));
  }
  return {
    data: new Float32Array(data),
    elementType: gl.LINE_STRIP,
    elementCount: this.numberOfVertices,
  };
};


Display.prototype.debugView = function(gl) {
  var theta = Math.PI / 180.0 * this.angle;
  var cos0 = Math.cos(theta);
  var sin0 = Math.sin(theta);
  var data = [];
  for (var i = 0; i < 128; ++i) {
    var angle0 = 2.0 * Math.PI / 128.0 * (i + 0);
    var angle1 = 2.0 * Math.PI / 128.0 * (i + 1)
    data.push(0.5 * Math.cos(angle0), 0.0, 0.5 * Math.sin(angle0), 1.0, 1.0, 1.0, 1.0);
    data.push(0.5 * Math.cos(angle1), 0.0, 0.5 * Math.sin(angle1), 1.0, 1.0, 1.0, 1.0);
  }
  return {
    data: new Float32Array(data),
    elementType: gl.LINES,
    elementCount: 256
  }
};


Display.prototype.poll = function() {
  var result = {
    b: this.currentState,
    c: this.lastError
  };
  this.lastError = Display.Error.NONE;
  return result;
};


Display.prototype.map = function(x, y) {
  this.memoryMapOffset = x;
  this.numberOfVertices = y;
  this.lastError = this.numberOfVertices > 128 ? Display.Error.BROKEN : Display.Error.NONE;
};


Display.prototype.rotateTo = function(x) {
  this.targetAngle = x;
};


Display.prototype.lerp = function(a, b, t) {
  return (1.0 - t) * a + t * b;
};


Display.prototype.update = function(dt) {
  this.currentState = this.numberOfVertices ? Display.State.RUNNING : Display.State.NO_DATA;
  var angle = this.lerp(this.angle, this.targetAngle, Display.LERP_ALPHA);
  this.currentState = angle - this.angle > Display.TURN_EPSILON ?
      Display.State.TURNING : this.currentState;
  this.angle = angle;
};


Display.prototype.view = function(gl) {
  var data = [];
  var count = 0;
  var alpha = 1.0 / Math.sqrt(this.numberOfVertices) / 2.0;
  var offset = -4.0 * Math.PI / 3.0 + Math.PI / 2.0;
  var theta0 = 0.0 + offset;
  var theta1 = 2.0 * Math.PI / 3.0 + offset;
  var theta2 = 4.0 * Math.PI / 3.0 + offset;
  var origin = {
    4: [0.5 * Math.cos(theta0), 0.0, 0.5 * Math.sin(theta0), 0.0, 0.0, 0.0, alpha],
    2: [0.5 * Math.cos(theta1), 0.0, 0.5 * Math.sin(theta1), 0.0, 0.0, 0.0, alpha],
    1: [0.5 * Math.cos(theta2), 0.0, 0.5 * Math.sin(theta2), 0.0, 0.0, 0.0, alpha]
  };
  for (var i = 0; i < this.numberOfVertices - 1; ++i) {
    var word1 = this.dcpu.memory[2 * i + 1];
    var word3 = this.dcpu.memory[2 * i + 3];
    var color0 = (word1 >> Display.COLOR_SHIFT) & Display.COLOR_MASK;
    var color1 = (word3 >> Display.COLOR_SHIFT) & Display.COLOR_MASK;
    var v0 = this.buildVertex(this.dcpu.memory[2 * i + 0], this.dcpu.memory[2 * i + 1], alpha);
    var v1 = this.buildVertex(this.dcpu.memory[2 * i + 2], this.dcpu.memory[2 * i + 3], alpha);
    var v0Black = [v0[0], v0[1], v0[2], 0.0, 0.0, 0.0, v0[6]];
    var v1Black = [v1[0], v1[1], v1[2], 0.0, 0.0, 0.0, v1[6]];
    for (var c = 1; c <= 4; c <<= 1) {
      if ((color0 & c) > 0 || (color1 & c) > 0) {
        count += 1;
        var v0Prime = [v0[0], v0[1], v0[2], 4 == c, 2 == c, 1 == c, v0[6]];
        var v1Prime = [v1[0], v1[1], v1[2], 4 == c, 2 == c, 1 == c, v1[6]];
        Array.prototype.push.apply(data, (color0 & c) > 0 ? v0Prime : v0Black);
        Array.prototype.push.apply(data, (color1 & c) > 0 ? v1Prime : v1Black);
        Array.prototype.push.apply(data, origin[c]);
      }
    }
  }
  return {
    data: new Float32Array(data),
    elementType: gl.TRIANGLES,
    elementCount: 3 * count
  };
};

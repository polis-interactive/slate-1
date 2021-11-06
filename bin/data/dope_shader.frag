
#ifdef GL_ES

precision highp float;
#define IN varying
#define OUT out
#define TEXTURE texture2D

#else

#define IN in
#define OUT out
#define TEXTURE texture

#endif

uniform float time;
uniform vec2 resolution;

uniform float gamma;
uniform float speed;
uniform float scale;
uniform float brightness;
uniform float contrast;

#define TAU 6.28318530718
#define PI 3.14158

// main
void main(void) {

    float t1 = time;
    // uv should be the 0-1 uv of texture...
    vec2 uv = gl_FragCoord.xy / resolution.xy;

    vec3 color = vec3((sin(t1 + uv.y) + 1.0)*0.5);
    
    gl_FragColor = vec4(pow(color, vec3(gamma)), 1.0);

}

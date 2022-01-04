
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
uniform float speed;

void main(void)
{

    float t = time * speed;

    vec2 size = vec2(800, 600);

    vec2 uv = gl_FragCoord.xy / size.xy;

    gl_FragColor = vec4(uv.x, uv.y, sin(t), 1.0);
}
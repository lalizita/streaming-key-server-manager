ffmpeg -i livetopzera_1d4dbbd5-6b3f-4691-b4c6-9b62ad7e7f70.m3u8 -o livetopzera.m3u8 -map 0:v -c copy -map 0:a -c copy -segment_time 10 -segment_name "livetopzera_%d.ts"

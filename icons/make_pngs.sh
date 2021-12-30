# Simplified from this script:
# https://gist.github.com/3v1n0/bb991e635100220fd307ebab5310b968

RES=../resources

for file in svgs/*.svg; do
    base=$(basename "$file")
    png="$RES/${base%.svg}.png"
    rsvg-convert "$file" -w 48 -h 48 -f png -o "$png.tmp.png"
    convert "$png.tmp.png" \
        -channel RGB \
        -colorspace gray \
        -background black \
        -alpha remove \
        -alpha off \
        "$png"
    rm "$png.tmp.png"
done
convert $RES/right.png -rotate 90 $RES/down.png
convert $RES/right.png -rotate 180 $RES/left.png
convert $RES/right.png -rotate 270 $RES/up.png
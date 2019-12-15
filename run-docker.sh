    docker run -it --rm \
        --name=izi \
        -e "ENV=production" \
        -e "LOG_LEVEL=INFO" \
        -e "MONGERAL_AEGON_HOST=https://gateway.gr1d.io/production/mongeralaegon/v1/" \
        -e "COMPLINE_HOST=https://gateway.gr1d.io/production/compline/signature/v1/" \
        -e "BIG_ID_HOST=https://gateway.gr1d.io/production/bigdata/bigid/ocr/v1/" \
        -e "SERVER_PORT=8107"
        JhonasMutton/izzi
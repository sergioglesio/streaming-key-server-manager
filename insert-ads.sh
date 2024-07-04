#!/bin/bash

# Diretório onde os arquivos de anúncios estão localizados
ADS_DIR="/hls/ads"

# Diretório onde os segmentos HLS estão sendo armazenados
HLS_DIR="/hls/live"

# Playlist principal HLS
MAIN_PLAYLIST="$HLS_DIR/rirguns_b9087751262144e44b4d775d0506ddda.m3u8"

# Verifica se o arquivo de playlist de anúncios existe
if [ ! -f "$ADS_DIR/ads-heineken.m3u8" ]; then
    echo "Arquivo de playlist de anúncios não encontrado!"
    exit 1
fi

# Copia os segmentos de anúncios para o diretório HLS
echo "Copiando segmentos de anúncios para o diretório HLS..."
cp "$ADS_DIR"/*.ts "$HLS_DIR/"

# Backup da playlist principal
cp "$MAIN_PLAYLIST" "$MAIN_PLAYLIST.bak"

# Adiciona os segmentos de anúncio à playlist principal
echo "Adicionando anúncios à playlist principal..."
while IFS= read -r line
do
  if [[ $line == *.ts ]]; then
    echo "#EXTINF:10.0," >> "$MAIN_PLAYLIST"
    echo "$line" >> "$MAIN_PLAYLIST"
  fi
done < "$ADS_DIR/ads-heineken.m3u8"

# Verifica o conteúdo da playlist principal após a inserção dos anúncios
echo "Conteúdo da playlist principal após a inserção dos anúncios:"
cat "$MAIN_PLAYLIST"

# Verifica se os arquivos foram copiados
echo "ADS no diretório HLS"
ls -lh "$HLS_DIR"

echo "Anúncio inserido."

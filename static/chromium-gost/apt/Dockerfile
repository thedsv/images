FROM browsers/base:7.4.2

ARG VERSION
ARG PACKAGE=chromium-gost-stable
ARG INSTALL_DIR=chromium-gost

LABEL browser=$PACKAGE:$VERSION

RUN \
        curl -s https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
        echo 'deb [arch=amd64] https://dl.google.com/linux/chrome/deb/ stable main' > /etc/apt/sources.list.d/google.list && \
        apt-get update && \
        apt-get -y --no-install-recommends install iproute2 $PACKAGE=$VERSION && \
         ( \
          sed -i -e 's@exec -a "$0" "$HERE/chrome"@& --no-sandbox --disable-gpu@' /opt/$INSTALL_DIR/$PACKAGE || \
          sed -i -e 's@exec -a "$0" "$HERE/chrome"@& --no-sandbox --disable-gpu@' /opt/$INSTALL_DIR/chromium-gost \
         ) && \
        chown root:root /opt/$INSTALL_DIR/chrome-sandbox && \
        chmod 4755 /opt/$INSTALL_DIR/chrome-sandbox && \
        chromium-gost --version && \
        rm -Rf /tmp/* && rm -Rf /var/lib/apt/lists/*

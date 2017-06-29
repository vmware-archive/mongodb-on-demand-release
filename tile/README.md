# How to build MongoDB On Demand for PCF using tile generator

0. Make sure you have [tile-generator](https://github.com/cf-platform-eng/tile-generator) and [bosh_cli](https://bosh.io/docs/bosh-cli.html) installed.

1. Check out PCF MongoDB On Demand tile generator repo:

    ```bash
    git clone https://github.com/Altoros/mongodb-on-demand-release.git
    ```

2. Download [Pivotal Cloud Foundry On Demand Service Broker Release](https://network.pivotal.io/products/on-demand-services-sdk/) into [tile/resources](https://github.com/Altoros/mongodb-on-demand-release/tree/master/tile/resources) folder.

3. Download missing golang and libsnmp packages. 

    ```bash
    cd mongodb-on-demand-release
    mkdir -p src/golang src/libsnmp
    wget 'https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz' -P src/golang
    wget 'http://security.ubuntu.com/ubuntu/pool/main/p/perl/libperl5.18_5.18.2-2ubuntu1.1_amd64.deb' -P src/libsnmp
    wget 'http://security.ubuntu.com/ubuntu/pool/main/n/net-snmp/libsnmp-base_5.7.2~dfsg-8.1ubuntu3.1_all.deb' -P src/libsnmp
    wget 'http://security.ubuntu.com/ubuntu/pool/main/n/net-snmp/libsnmp30_5.7.2~dfsg-8.1ubuntu3.1_amd64.deb' -P src/libsnmp
    ```

4. Create release tarball and place it into [tile/resources](https://github.com/Altoros/mongodb-on-demand-release/tree/master/tile/resources) folder.

    ```bash
    cd mongodb-on-demand-release
    bosh create release --with-tarball --version $VERSION_NUMBER --force --name mongodb
    cp dev_releases/mongodb/mongodb-$VERSION_NUMBER.tgz tile/resources
    ```

5. Edit `tile.yml` file and check path and versions for mongodb-on-demand-release.
   Ensure that tile file configured with version which was specified in step 4. [#1](https://github.com/Altoros/mongodb-on-demand-release/blob/904e54b8998f32a8594971fb9a83255486d22dd2/tile/tile.yml#L66) [#2](https://github.com/Altoros/mongodb-on-demand-release/blob/904e54b8998f32a8594971fb9a83255486d22dd2/tile/tile.yml#L109)

    ```bash
    cd mongodb-on-demand-release/tile
    vi tile.yml
    ```

6. Build your tile, after build is finished you can find product file in `tile/product` subdirectory.

    ```bash
    tile build $VERSION_NUMBER
    ```

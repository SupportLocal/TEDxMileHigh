
TEDxMileHigh
============

Development
-----------

```bash
# fork it
# clone it
cd TEDxMileHigh
cp TEDxMileHigh.toml.example TEDxMileHigh.toml
$EDITOR TEDxMileHigh.toml # dial config in for your environment
bundle
rake develop
```

If everything goes well, guard is up and running now. Project changes
will build, test, vet and restart automatically. LiveReload is also
running, so web facing changes should be especially easy to monitor by
browsing http://localhost:9000 (assuming you stuck with :9000 in the
config).

Deployment
----------

Initial deployment looks like this (assuming you've already created and
locked down your server).

```bash
# from a clean workspace
rake clean package
scp pkg/TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2 <yourhost>:/opt/
ssh <yourhost>
cd /opt
tar xjf TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2
ln -s TEDxMileHigh-<your-branch>@<your-revision> TEDxMileHigh
cp TEDxMileHigh/etc/TEDxMileHigh.toml.example /etc/TEDxMileHigh.toml
$EDITOR /etc/TEDxMileHigh.toml # adjust for your environment
cp -r TEDxMileHigh/etc/init/*.conf /etc/init/
start TEDxMileHigh-boot
```

Ongoing deployments should look more like this:
```bash
# from a clean workspace
rake clean package
scp pkg/TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2 <yourhost>:/opt/
ssh <yourhost>
cd /opt
tar xjf TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2
rm TEDxMileHigh
ln -s TEDxMileHigh-<your-branch>@<your-revision> TEDxMileHigh
restart TEDxMileHigh-all
```

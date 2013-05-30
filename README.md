TEDxMileHigh
============


Development
-----------

fork it
clone your fork
cd TEDxMileHigh
cp TEDxMileHigh.toml.example TEDxMileHigh.toml
$EDITOR TEDxMileHigh.toml # dial config in for your environment
bundle
rake develop

If everything goes well, guard is up and running now. Project changes
will build, test, vet and restart automatically. LiveReload is also
running, so web facing changes should be especially easy to monitor by
browsing http://localhost:9000 (assuming you stuck with :9000 in the
config).


Deployment
----------

From a clean workspace, deployment looks like this:
Create a server.
Lock it down.

Initial deployments look like roughly like this:
rake clean package
scp pkg/TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2 <yourhost>:/opt/
ssh <yourhost>
cd /opt
tar xjf TEDxMileHigh-<your-branch>@<your-revision>.tar.bz2
ln -s TEDxMileHigh-<your-branch>@<your-revision> TEDxMileHigh


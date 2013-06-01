require 'rake/packagetask'

{
  'jslint' => 'npm install jslint -g',
  'mongo'  => 'brew install mongo',
  'go'     => 'brew install go --cross-compile-common',
}.each do |cmd, installer|
  task develop: "which:#{cmd}"

  desc "locates #{cmd} on your PATH or complains loudly"
  task "which:#{cmd}" do
    system("which #{cmd} > /dev/null") || fail("#{cmd} not found!\nHave you installed it?\nTry: `#{installer}`")
  end
end


%w(
  github.com/antage/eventsource/http
  github.com/darkhelmet/twitterstream
  github.com/garyburd/redigo/redis
  github.com/gorilla/mux
  github.com/laurent22/toml-go/toml
).each do |dep|
  dep_path = File.join(ENV['GOPATH'], dep)

  file dep_path do
    system("go get #{dep}") || fail
  end

  task 'go:dependencies' => dep_path
end


file TEDxMileHigh: 'go:dependencies' do
  system('go build') || fail
end


namespace :go do
	desc 'go get declared dependencies'
	task :dependencies
end


desc 'Start developing!'
task develop: :TEDxMileHigh do
  system('bundle') || fail
  exec('bundle exec guard')
end

task default: :develop


# ----------- packaging for distribution

task 'clean' do
  rm_rf 'bin'
  rm_rf 'pkg'
end

task 'ensure-clean-workspace' do
  unless true # `git status --porcelain`.chomp == ''
    puts "your workspace isn't clean. exiting"
    exit(1)
  end
end

task package: 'ensure-clean-workspace' # fail-safe; don't package source tree with unmanaged and/or uncommitted files


$css_map = Dir['assets/sass/*.scss'].inject(Hash.new) do |css_map, scss_file|
  css_file = scss_file.gsub(%r{/sass/}, '/css/').gsub(%r{scss$}, 'css')
  css_map[css_file] = scss_file
  css_map
end


$go_files = Dir['**/*.go']
$binaries = []

[ ['linux',  ['amd64']],
  ['darwin', ['amd64']],
].each do |goos, goarchs|

  goarchs.each do |goarch|

    arch_dir = "bin/#{goos}/#{goarch}"
    directory arch_dir => 'ensure-clean-workspace'

    binary = "bin/#{goos}/#{goarch}/TEDxMileHigh"
    file binary => [arch_dir] + $go_files do
      begin
        ENV["GOOS"] = goos
        ENV["GOARCH"] = goarch
        ENV["CGO_ENABLED"] = "0"
        system(%Q(go build -o "#{binary}")) || fail

      ensure
        ENV.delete("GOOS")
        ENV.delete("GOARCH")
        ENV.delete("CGO_ENABLED")
      end
    end

    $binaries.push(binary)
  end
end

$branch = `git rev-parse --abbrev-ref HEAD`.chomp
$revision = `git rev-parse --short HEAD`.chomp

package = Rake::PackageTask.new('TEDxMileHigh', "#{$branch}@#{$revision}") do |pt|
  pt.need_tar_bz2 = true

  $css_map.each do |css_file, scss_file|
    pt.package_files.include(css_file)
    pt.package_files.exclude(scss_file)
  end

  $binaries.each do |binary|
    pt.package_files.include(binary)
  end

  pt.package_files.include 'assets/ejs/**/*.ejs'
  pt.package_files.include 'assets/img/**/*.{gif,jpg,jpeg,png,svg}'
  pt.package_files.include 'assets/js/**/*.js'
  pt.package_files.include 'assets/vendor/**/*.*'
  pt.package_files.include 'etc/TEDxMileHigh.toml.example'
  pt.package_files.include 'etc/init/*.conf'
end

$archive_file = File.join(package.package_dir, package.tar_bz2_file)

file $archive_file => $binaries

$css_map.each do |css_file, scss_file|
  task "clean:#{css_file}" => 'ensure-clean-workspace' do
    rm_f css_file
  end

  file css_file => "clean:#{css_file}" do
    system %Q(compass compile --config=assets/config.rb --output-style=compressed --boring --quiet "#{scss_file}") || fail
  end

  file $archive_file => css_file
end

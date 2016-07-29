VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "debian-8.1.0"
  config.vm.box_url = "https://github.com/kraksoft/vagrant-box-debian/releases/download/8.1.0/debian-8.1.0-amd64.box"
  config.vm.network "private_network", ip: "192.168.10.110"
  config.vm.synced_folder ".", "/vagrant", disabled: true
end


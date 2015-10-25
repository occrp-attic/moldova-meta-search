Vagrant.require_version ">= 1.5"

# Check to determine whether we're on a windows or linux/os-x host,
# later on we use this to launch ansible in the supported way
# source: https://stackoverflow.com/questions/2108727/which-in-ruby-checking-if-program-exists-in-path-from-ruby
def which(cmd)
    exts = ENV['PATHEXT'] ? ENV['PATHEXT'].split(';') : ['']
    ENV['PATH'].split(File::PATH_SEPARATOR).each do |path|
        exts.each { |ext|
            exe = File.join(path, "#{cmd}#{ext}")
            return exe if File.executable? exe
        }
    end
    return nil
end

Vagrant.configure("2") do |config|
    config.vm.provider :virtualbox do |v|
        v.name = "vagrant-ansible"
		
        v.customize [
            "modifyvm", :id,
            "--name", "vagrant-ansible",
            "--memory", 1024,
            "--natdnshostresolver1", "on",
            "--cpus", 4,
        ]
    end

    config.vm.box = "ubuntu/trusty64"

    config.vm.network :private_network, ip: "192.168.33.110"

    #config.ssh.forward_agent = true

    #############################################################
    # Ansible provisioning (you need to have ansible installed)
    #############################################################

    if which('ansible-playbook')
        config.vm.provision "ansible" do |ansible|
            ansible.playbook = "ansible/playbook.yml"
            ansible.inventory_path = "ansible/inventories/dev"
            ansible.limit = 'all'
        end

		config.vm.synced_folder "./", "/vagrant", type: "nfs"
    else
	
		# fix for windows - long path and symlinks
		# https://gist.github.com/Jakobud/0768ff6b6051e79eef60

		# First, disable default vagrant share
		config.vm.synced_folder ".", "/vagrant", disabled: true

		# Next, setup the shared Vagrant folder manually, bypassing Windows 260 character path limit
		config.vm.provider "virtualbox" do |v|
			v.customize ["sharedfolder", "add", :id, "--name", "vagrant", "--hostpath", (("//?/" + File.dirname(__FILE__)).gsub("/","\\"))]
			v.customize ["setextradata", :id, "VBoxInternal2/SharedFoldersEnableSymlinksCreate/vagrant", "1"]
		end

		# Finally, mount the shared folder on the guest system during provision
		config.vm.provision :shell, inline: "mkdir -p /vagrant", run: "always"
		config.vm.provision :shell, inline: "mount -t vboxsf -o uid=`id -u vagrant`,gid=`getent group vagrant | cut -d: -f3` vagrant /vagrant", run: "always"

	
        config.vm.provision :shell, path: "ansible/windows.sh", args: ["vagrant-ansible"]
    end
end

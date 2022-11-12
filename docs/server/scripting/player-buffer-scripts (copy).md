# Perl Plugins

Inside the [plugins](https://github.com/ProjectEQ/projecteqquests/tree/master/plugins) directory you can find example scripts on how to create reusable methods so that you don't have to copy and paste code in between files.

The `plugins` folder can be located in the `quests` folder and can be either dragged to the base folder directory or [symbollically linked](https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mklink) to be called from within the `quests` folder.

An example plugin is this:
```pl
sub CustomData {
	my %custom_data = (
		0 => "A",
		1 => "B",
		2 => "C",
		3 => "D"
	);

	return %custom_data;
}

sub ListCustomData {
	my %custom_data = CustomData();

	foreach my $key (sort {$a <=> $b} keys %custom_data) {
		quest::message(315, "$key: $custom_data{$key}");
	}
}
```

You can call this `custom_plugin.pl`.

An example of an NPC using this would be as follows.
```pl
sub EVENT_SAY {
	if ($text=~/Hail/i) {
		plugin::ListCustomData();
	}
}
```

![image](https://user-images.githubusercontent.com/89047260/201449965-c38e95fd-e1e4-49dd-ab48-e4434af3a8d1.png)

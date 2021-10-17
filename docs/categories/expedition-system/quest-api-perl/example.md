# Example

This is a full example to demonstrate usage. Guard Pineshade at the orc lift in gfaydark provides an expedition to crushbone

`gfaydark/Guard_Pineshade.pl`

```perl
my $expedition_name = "Crushbone, DVinn's Stronghold";
my $min_players     = 1;
my $max_players     = 6;
my $dz_zone         = "crushbone";
my $dz_version      = 0;
my $dz_duration     = 3600; # 1 hour

sub EVENT_SAY {
  if ($text =~ /hail/i) {
    my $dz = $client->GetExpedition();
    if ($dz && $dz->GetName() eq $expedition_name) {
      quest::say("Tell me when you're [" . quest::saylink("ready") . "] to enter");
    }
    else {
      quest::say("Would you like to [" . quest::saylink("request") . "] the expedition?");
    }
  }
  elsif ($text =~ /request/i) {
    my $dz = $client->CreateExpedition($dz_zone, $dz_version, $dz_duration, $expedition_name, $min_players, $max_players);
    if ($dz) {
      $dz->SetCompass("gfaydark", 238.0, 987.0, -24.90); # pointing to guard pineshade
      $dz->SetSafeReturn("gfaydark", 245.84, 987.93, -27.6, 484.0); # orc lift in gfay
      $dz->SetZoneInLocation(479.44, -500.18, 5.75, 421.8); # bridge in crushbone
      $dz->AddReplayLockout(7200); # immediately add a 2 hour replay lockout on creation
      quest::say("Tell me when you're [" . quest::saylink("ready") . "] to enter");
    }
  }
  elsif ($text =~ /ready/i) {
    $client->MovePCDynamicZone("crushbone");
  }
}
```

`crushbone/zone_controller.pl`

```perl
# this requires the Zone:UseZoneController rule to be enabled
my $initialized = 0;

sub Initialize {
  quest::debug("Initializing expedition");
  my $dz = quest::get_expedition();
  if ($dz) {
    # bind unique npc types to events to prevent any looting exploits
    $dz->SetLootEventByNPCTypeID(58032, "Emperor Crush"); # npc: Emperor_Crush
    $dz->SetLootEventByNPCTypeID(317109, "Ambassador Mata Muram"); # npc: Overlord_Mata_Muram

    # spawn any events that the expedition doesn't have a lockout for
    if (!$dz->HasLockout("Ambassador Mata Muram")) {
      quest::spawn2(317109, 0, 0, 130.78, -149.09, 88.70, 270.8); # npc: Overlord_Mata_Muram
    }
  }
}

sub EVENT_SPAWN_ZONE {
  # note: zone_controller spawns before expeditions can be cached, so we delay
  # expedition init until first npc spawns instead of using our EVENT_SPAWN
  if (!$initialized) {
    $initialized = 1; # set first to prevent infinite recursion from spawning npcs in init()
    Initialize();
  }

  if ($spawned_npc_id == 58032) { # npc: Emperor_Crush
    my $dz = quest::get_expedition();
    if ($dz && $dz->HasLockout("Emperor Crush")) {
      quest::depop(58032); # depop emperor crush due to lockout
    }
  }
}

sub EVENT_DEATH_ZONE {
  if ($killed_npc_id == 58032) { # npc: Emperor_Crush
    my $dz = quest::get_expedition();
    if ($dz) {
      $dz->AddLockout("Emperor Crush", 86400); # 1 day lockout
    }
  }
  elsif ($killed_npc_id == 317109) { # npc: Overlord_Mata_Muram
    my $dz = quest::get_expedition();
    if ($dz) {
      $dz->AddLockout("Ambassador Mata Muram", 604800); # 7 day lockout

      # spawn a chest and bind its spawn id to the event to prevent loot exploits
      my $chest_id = quest::spawn2(893, 0, 0, 130.78, -149.09, 88.70, 0);
      $dz->SetLootEventBySpawnID($chest_id, "Ambassador Mata Muram");
    }
  }
}
```


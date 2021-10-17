# EQDictionary

#### What is it?

EQDictionary is a collection of lookup references, accessed by specific namespaces and referenced by implementation standard, or [client](https://github.com/EQEmu/Server/wiki/Client-Version-List) and [mob](https://github.com/EQEmu/Server/wiki/Mob-Version-List) versions.

References are usually tied back to the originating client definition and are available through an indexed lookup system. Server-based 'hybrid' definitions may also be implemented to bridge differences between client implementations and server requirements.

The purpose of this system is to allow cyclic or minimum conditional evaluations in order to transform data.\


#### What can you do with it?

Simplifying coding methodology can change current standards from this:

```
if (ClientVersion() >= EQEmu::versions::ClientVersion::SoF) { bank_slot_count = 24; }

else { bank_slot_count = 16; }

if (current_slot >= bank_slot_count) { return; }
```

Or this:

```
if (m_ClientBitmask & EQEmu::versions::maskSoFAndLater) { bank_slot_count = 24; }

else { bank_slot_count = 16; }

if (current_slot >= bank_slot_count) { return; }
```

To this:

```
if (current_slot >= m_inv.GetInv().GetLookup()->InventoryTypeSize.Bank) { return; }
```

The first case requires that clients to be added in order of release.

The second case will allow clients to be added in any order..but, requires awkward manipulation of mask bits.

The EQDictionary (last) case does not require any special ordering of clients or manipulation of mask bits. In fact, this system makes ClientVersion comparison and bitmask use obsolete.

Additionally, duplicate definitions can be avoided by having a singular location to reference.\


#### What makes up EQDictionary?

Static and Dynamic references.

Static references are design-/compile-time definitions of maximum client version-based values and are linked in 2 ways:

* Server implementation standard accessed through `EQEmu::<system_namespace>::<property>`
* Server-based version lookup accessed through `EQEmu::<system_namespace>::Lookup(version)-><property>`

_Note: invalid versions return the default, null-set reference_

Dynamic references are run-time definitions based on expansion settings, and possibly other criteria.

There are 2 dynamic references based on gm flag states `<set>` and `<clear>`.

Dynamic lookups are usually embedded into a system class..but, their use is not exclusive to them.

Not every case of client/mob version and gm flag set/clear will receive a dynamic entry. In these cases, the static reference is returned.

The entries are constructed based on expansion settings when the initial call to `EQEmu::<system_namespace>::InitializeDynamicLookups()` is performed at server start-up.

Creating dynamic lookups and embedding them into system classes allow for criteria updates that would normally require even more state condition checks and more static variable definitions making proper support for per-expansion and gm flag states without them impractical.

(The number of static lookups required to support non-dynamic lookups would essentially be 'the number of expansions' -> squared, then multiplied by 2 to account for gm state.)

Dynamic lookups resolve the complexity of coding needed to support the myriad of conditions required support to customizable server features.

A deep, working knowledge of client behavior is required to correctly set up dynamic entries. Testing through server setting changes and client observation is the best way to achieve this.\


#### Where do I look for all of this?

Currently, the following files comprise the system:

* [emu_constants.h](https://github.com/EQEmu/Server/blob/master/common/emu_constants.h)
* [emu_limits.h](https://github.com/EQEmu/Server/blob/master/common/emu_limits.h)
* [emu_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/emu_limits.cpp)
* [eq_limits.h](https://github.com/EQEmu/Server/blob/master/common/eq_limits.h)
* [eq_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/eq_limits.cpp)
* [titanium_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/titanium_limits.h)
* [titanium_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/titanium_limits.cpp)
* [sof_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/sof_limits.h)
* [sof_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/sof_limits.cpp)
* [sod_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/sod_limits.h)
* [sod_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/sod_limits.cpp)
* [uf_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/uf_limits.h)
* [uf_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/uf_limits.cpp)
* [rof_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/rof_limits.h)
* [rof_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/rof_limits.cpp)
* [rof2\_limits.h](https://github.com/EQEmu/Server/blob/master/common/patches/rof2\_limits.h)
* [rof2\_limits.cpp](https://github.com/EQEmu/Server/blob/master/common/patches/rof2\_limits.cpp)

#### How do I create a new lookup or add to an existing one?

Adding properties to an existing lookup involves most of the same steps as creating a new lookup system.

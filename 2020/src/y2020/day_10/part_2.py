DEBUG = True

# ===================================== SOLUTION =====================================
# Only adapters with a distance under 3 can have different arrangements
# For a given adapter, the combinations to reach it are given by the
# sum of combinations of adapters with distance of 1, 2 or 3, if they exist
# For example, given:
# 1 - 2 - 3 - 5
# We first add 0 at the beginning of the adapters chain, and the adapter 8 at the end
# 0 - 1 - 2 - 3 - 5 - 8
# We can reach 0 in 1 way
# We store this value in a dictionary so that we can access it in ~ constant time (O(1))
# {0: 1}
# We move to 1. The only combination to reach it is through 0.
# {0: 1, 1: 1}
# We move to 2. We can branch 2 in 1, or directly in 0 as their distance is under 3
# This is the sum of possibilities to reach previous 2 adapters
# {0: 1, 1: 1, 2: 2}
# We move to 3. We can branch 3 in 2, 1, or 0: the possibilities to branch it are the sum
# of possibilities to reach those 3 adapters
# {0: 1, 1: 1, 2: 2, 3: 4}
# We move to 5. We can branch 5 in 3 or 2 as we have no adapter 4. So to reach 5 we have
# possibilities of adapter 3 plus possibilities of adapter 2, that gives 6.
# {0: 1, 1: 1, 2: 2, 3: 4, 5: 6}
# We move to last adapter, 8, and it can only be branched into adapter 5, adding no more
# combination.
# {0: 1, 1: 1, 2: 2, 3: 4, 5: 6, 8: 6}
# The result is the combinations to reach last adapter: 6


def dbg_print(s):
    if DEBUG:
        print(s)


def solution(data):
    adapters = [int(x) for x in data.splitlines()]
    adapters.append(0)
    adapters.sort()
    adapters.append(adapters[-1] + 3)
    possibilities_to_reach_adapter_map = {0: 1}
    for adapter in adapters[1:]:
        possibilities_to_reach_adapter_map[adapter] = (
            possibilities_to_reach_adapter_map.get(adapter - 1, 0)
            + possibilities_to_reach_adapter_map.get(adapter - 2, 0)
            + possibilities_to_reach_adapter_map.get(adapter - 3, 0)
        )
    return possibilities_to_reach_adapter_map[adapters[-1]]

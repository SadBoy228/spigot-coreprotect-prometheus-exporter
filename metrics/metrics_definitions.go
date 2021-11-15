package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    ServerPlayerWheneverConnected = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_server_player_connection_counter",
        Help: "Counter, that represents quantity of players, whenever connected to server",
    })

    ServerPlayerCurrentOnline = prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Name: "cp_prometheus_exporter_server_player_current_online",
        Help: "Gauge, that represents quantity of players, that currently online on server",
    },
        []string{"world"},
    )

    PlaytimeAmount = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_playtime_amount",
        Help: "Counter, that represents overall time, that players spent during playing on server",
    },
        []string{"player"},
    )

    PlaytimeAverage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Name: "cp_prometheus_exporter_playtime_average",
        Help: "Gauge, that represents averate time players spent during ordinary player session",
    },
        []string{"player"},
    )

    ChatAmount = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_chat_amount",
        Help: "Counter, that represent overall quantity of chat messages, that players sent during playing on server",
    },
        []string{"player"},
    )

    ChatAverage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Name: "cp_prometheus_exporter_chat_average",
        Help: "Gauge, that represents average quantity of chat messages players sent during ordinary player session",
    },
        []string{"player"},
    )

    CommandUsageAmount = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_command_usage_amount",
        Help: "Counter, that represents amount of command usage during overall time of server work",
    },
        []string{"command", "player"},
    )

    CommandUsageAverage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Name: "cp_prometheus_exporter_command_usage_average",
        Help: "Gauge, that represents average command usage during ordinary player session",
    },
        []string{"command", "player"},
    )

    WorldBlockPermutationClaimed = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_world_block_permutation_claimed",
        Help: "Counter, that represents overall world block claim permutations by players, block IDs and world, where permutation happened",
    },
        []string{"player", "block", "world"},
    )

    WorldBlockPermutationBroke = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_world_block_permutation_broke",
        Help: "Counter, that represents overall world block break permurations by players, block IDs and world, where permutation happened",
    },
        []string{"player", "block", "world"},
    )

    WorldItemPermutationClaimed = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_world_item_permutation_claimed",
        Help: "Counter, that represents overall world item claim permutations by players, item IDs and world, where permutation happened",
    },
        []string{"player", "item", "world"},
    )

    WorldItemPermutationStorageDeposited = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_world_item_permutation_storage_deposited",
        Help: "Counter, that represents overall world item storage deposit permutations by players, item IDs and world, where permutation happened",
    },
        []string{"player", "item", "world"},
    )

    WorldItemPermutationStorageWithdrew = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_world_item_permutation_storage_withdrew",
        Help: "Counter, that represents overall world item storage withdraw permutations by players, item IDs and world, where permutation happened",
    },
        []string{"player", "item", "world"},
    )

    PlayerDeathAmount = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "cp_prometheus_exporter_player_death_amount",
        Help: "Counter, that represents overall amount of players death by players (only death that caused by another player currently)",
    },
        []string{"player"},
    )
)

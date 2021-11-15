package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    packageRegistry *prometheus.Registry
)

func GetMetricRegistry() (*prometheus.Registry, error) {
    if packageRegistry == nil {
        newRegistry := prometheus.NewPedanticRegistry()

        collectorsList := []prometheus.Collector{
            ServerPlayerWheneverConnected,
            ServerPlayerCurrentOnline,
            PlaytimeAmount,
            PlaytimeAverage,
            ChatAmount,
            ChatAverage,
            CommandUsageAmount,
            CommandUsageAverage,
            WorldBlockPermutationClaimed,
            WorldBlockPermutationBroke,
            WorldItemPermutationClaimed,
            WorldItemPermutationStorageDeposited,
            WorldItemPermutationStorageWithdrew,
            PlayerDeathAmount,
        }

        for _, collector := range collectorsList {
            if err := newRegistry.Register(collector); err != nil {
                return nil, err
            }
        }

        packageRegistry = newRegistry
    }

    return packageRegistry, nil
}

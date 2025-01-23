package config

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap/zapcore"

	eotscfg "github.com/Manta-Network/manta-fp/eotsmanager/config"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/util"
)

const (
	defaultChainType                   = "babylon"
	defaultLogLevel                    = zapcore.InfoLevel
	defaultLogDirname                  = "logs"
	defaultLogFilename                 = "fpd.log"
	defaultFinalityProviderKeyName     = "bbn-fp"
	DefaultRPCPort                     = 12581
	defaultConfigFileName              = "fpd.conf"
	defaultNumPubRand                  = 70000 // support running of 1 week with block production time as 10s
	defaultNumPubRandMax               = 100000
	defaultMinRandHeightGap            = 35000
	defaultBatchSubmissionSize         = 1000
	defaultStatusUpdateInterval        = 20 * time.Second
	defaultRandomInterval              = 30 * time.Second
	defaultSubmitRetryInterval         = 1 * time.Second
	defaultSignatureSubmissionInterval = 1 * time.Second
	defaultMaxSubmissionRetries        = 20
	defaultBitcoinNetwork              = "signet"
	defaultDataDirname                 = "data"
)

var (
	DefaultFpdDir             = btcutil.AppDataDir("fpd", false)
	defaultBTCNetParams       = chaincfg.SigNetParams
	defaultEOTSManagerAddress = "127.0.0.1:" + strconv.Itoa(eotscfg.DefaultRPCPort)
	DefaultRPCListener        = "127.0.0.1:" + strconv.Itoa(DefaultRPCPort)
	DefaultDataDir            = DataDir(DefaultFpdDir)
)

type Config struct {
	LogLevel                    string        `long:"loglevel" description:"Logging level for all subsystems" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal"`
	ChainType                   string        `long:"chaintype" description:"the type of the consumer chain" choice:"babylon"`
	NumPubRand                  uint32        `long:"numPubRand" description:"The number of Schnorr public randomness for each commitment"`
	NumPubRandMax               uint32        `long:"numpubrandmax" description:"The upper bound of the number of Schnorr public randomness for each commitment"`
	MinRandHeightGap            uint32        `long:"minrandheightgap" description:"The minimum gap between the last committed rand height and the current Babylon block height"`
	MaxSubmissionRetries        uint32        `long:"maxsubmissionretries" description:"The maximum number of retries to submit finality signature or public randomness"`
	EOTSManagerAddress          string        `long:"eotsmanageraddress" description:"The address of the remote EOTS manager; Empty if the EOTS manager is running locally"`
	BatchSubmissionSize         uint32        `long:"batchsubmissionsize" description:"The size of a batch in one submission"`
	StatusUpdateInterval        time.Duration `long:"statusupdateinterval" description:"The interval between each update of bbn-fp status"`
	RandomnessCommitInterval    time.Duration `long:"randomnesscommitinterval" description:"The interval between each attempt to commit public randomness"`
	SubmissionRetryInterval     time.Duration `long:"submissionretryinterval" description:"The interval between each attempt to submit finality signature or public randomness after a failure"`
	SignatureSubmissionInterval time.Duration `long:"signaturesubmissioninterval" description:"The interval between each finality signature(s) submission"`

	BitcoinNetwork string `long:"bitcoinnetwork" description:"Bitcoin network to run on" choise:"mainnet" choice:"regtest" choice:"testnet" choice:"simnet" choice:"signet"`

	BTCNetParams chaincfg.Params

	PollerConfig *ChainPollerConfig `group:"chainpollerconfig" namespace:"chainpollerconfig"`

	OpEventConfig *OpEventConfig `group:"opeventconfig" namespace:"opeventconfig"`

	DatabaseConfig *DBConfig `group:"dbconfig" namespace:"dbconfig"`

	BabylonConfig *BBNConfig `group:"babylon" namespace:"babylon"`

	RPCListener string `long:"rpclistener" description:"the listener for RPC connections, e.g., 127.0.0.1:1234"`

	Metrics *metrics.Config `group:"metrics" namespace:"metrics"`
}

func DefaultConfigWithHome(homePath string) Config {
	bbnCfg := DefaultBBNConfig()
	bbnCfg.Key = defaultFinalityProviderKeyName
	bbnCfg.KeyDirectory = homePath
	pollerCfg := DefaultChainPollerConfig()
	opEventConfig := DefaultOpEventConfig()
	cfg := Config{
		ChainType:                   defaultChainType,
		LogLevel:                    defaultLogLevel.String(),
		DatabaseConfig:              DefaultDBConfigWithHomePath(homePath),
		BabylonConfig:               &bbnCfg,
		PollerConfig:                &pollerCfg,
		OpEventConfig:               &opEventConfig,
		NumPubRand:                  defaultNumPubRand,
		NumPubRandMax:               defaultNumPubRandMax,
		MinRandHeightGap:            defaultMinRandHeightGap,
		BatchSubmissionSize:         defaultBatchSubmissionSize,
		StatusUpdateInterval:        defaultStatusUpdateInterval,
		RandomnessCommitInterval:    defaultRandomInterval,
		SubmissionRetryInterval:     defaultSubmitRetryInterval,
		SignatureSubmissionInterval: defaultSignatureSubmissionInterval,
		MaxSubmissionRetries:        defaultMaxSubmissionRetries,
		BitcoinNetwork:              defaultBitcoinNetwork,
		BTCNetParams:                defaultBTCNetParams,
		EOTSManagerAddress:          defaultEOTSManagerAddress,
		RPCListener:                 DefaultRPCListener,
		Metrics:                     metrics.DefaultFpConfig(),
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	return cfg
}

func DefaultConfig() Config {
	return DefaultConfigWithHome(DefaultFpdDir)
}

func CfgFile(homePath string) string {
	return filepath.Join(homePath, defaultConfigFileName)
}

func LogDir(homePath string) string {
	return filepath.Join(homePath, defaultLogDirname)
}

func LogFile(homePath string) string {
	return filepath.Join(LogDir(homePath), defaultLogFilename)
}

func DataDir(homePath string) string {
	return filepath.Join(homePath, defaultDataDirname)
}

func LoadConfig(homePath string) (*Config, error) {
	cfgFile := CfgFile(homePath)
	if !util.FileExists(cfgFile) {
		return nil, fmt.Errorf("specified config file does "+
			"not exist in %s", cfgFile)
	}

	var cfg Config
	fileParser := flags.NewParser(&cfg, flags.Default)
	err := flags.NewIniParser(fileParser).ParseFile(cfgFile)
	if err != nil {
		return nil, err
	}

	// Make sure everything we just loaded makes sense.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Validate() error {
	if cfg.EOTSManagerAddress == "" {
		return fmt.Errorf("EOTS manager address not specified")
	}

	btcNetConfig, err := NetParamsBTC(cfg.BitcoinNetwork)
	if err != nil {
		return err
	}
	cfg.BTCNetParams = btcNetConfig

	_, err = net.ResolveTCPAddr("tcp", cfg.RPCListener)
	if err != nil {
		return fmt.Errorf("invalid RPC listener address %s, %w", cfg.RPCListener, err)
	}

	if cfg.Metrics == nil {
		return fmt.Errorf("empty metrics config")
	}

	if err := cfg.Metrics.Validate(); err != nil {
		return fmt.Errorf("invalid metrics config")
	}

	return nil
}

func NetParamsBTC(btcNet string) (chaincfg.Params, error) {
	switch btcNet {
	case "mainnet":
		return chaincfg.MainNetParams, nil
	case "testnet":
		return chaincfg.TestNet3Params, nil
	case "regtest":
		return chaincfg.RegressionNetParams, nil
	case "simnet":
		return chaincfg.SimNetParams, nil
	case "signet":
		return chaincfg.SigNetParams, nil
	default:
		return chaincfg.Params{}, fmt.Errorf("invalid network: %v", btcNet)
	}
}

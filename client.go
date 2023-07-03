package cnc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	c *resty.Client
}

func NewClient(baseURL string, options ...Option) (*Client, error) {
	c := &Client{
		c: resty.New(),
	}

	c.c.SetBaseURL(baseURL)

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Header(ctx context.Context, height uint64) /* Header */ error {
	_ = headerPath()
	return errors.New("method Header not implemented")
}

func (c *Client) Balance(ctx context.Context) error {
	_ = balanceEndpoint
	return errors.New("method Balance not implemented")
}

func (c *Client) SubmitTx(ctx context.Context, tx []byte) /* TxResponse */ error {
	_ = submitTxEndpoint
	return errors.New("method SubmitTx not implemented")
}

func (c *Client) SubmitPFB(ctx context.Context, namespace Namespace, data []byte, fee int64, gasLimit uint64) (*TxResponse, error) {
	req := SubmitPFBRequest{
		NamespaceID: hex.EncodeToString(namespace.Bytes()),
		Data:        hex.EncodeToString(data),
		Fee:         fee,
		GasLimit:    gasLimit,
	}
	var res TxResponse
	var rpcErr string
	_, err := c.c.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		SetError(&rpcErr).
		Post(submitPFBEndpoint)
	if err != nil {
		return nil, err
	}
	if rpcErr != "" {
		return nil, errors.New(rpcErr)
	}

	var respErr error
	if res.Code != 0 {
		respErr = errors.Join(err, fmt.Errorf("codespace: %s, code: %v, err: %s", res.Codespace, res.Code, res.Logs.String()))
	}
	return &res, respErr
}

func (c *Client) NamespacedShares(ctx context.Context, namespace Namespace, height uint64) ([][]byte, error) {
	var res struct {
		Shares [][]byte `json:"shares"`
		Height uint64   `json:"height"`
	}

	err := c.callNamespacedEndpoint(ctx, namespace, height, namespacedSharesEndpoint, &res)
	if err != nil {
		return nil, err
	}

	return res.Shares, nil
}

func (c *Client) NamespacedData(ctx context.Context, namespace Namespace, height uint64) ([][]byte, error) {
	var res struct {
		Data   [][]byte `json:"data"`
		Height uint64   `json:"height"`
	}

	err := c.callNamespacedEndpoint(ctx, namespace, height, namespacedDataEndpoint, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

// callNamespacedEndpoint fetches result of /namespaced_{type} family of endpoints into result (this should be pointer!)
func (c *Client) callNamespacedEndpoint(ctx context.Context, namespace Namespace, height uint64, endpoint string, result interface{}) error {
	var rpcErr string
	_, err := c.c.R().
		SetContext(ctx).
		SetResult(result).
		SetError(&rpcErr).
		Get(namespacedPath(endpoint, namespace, height))
	if err != nil {
		return err
	}
	if rpcErr != "" {
		return errors.New(rpcErr)
	}
	return nil
}

func headerPath() string {
	return fmt.Sprintf("%s/%s", headerEndpoint, heightKey)
}

func namespacedPath(endpoint string, namespace Namespace, height uint64) string {
	return fmt.Sprintf("%s/%s/height/%d", endpoint, hex.EncodeToString(namespace.Bytes()), height)
}

package etcd

import (
	"io"

	"github.com/coreos/agro/models"
)

func (c *etcdCtx) DumpMetadata(w io.Writer) error {
	io.WriteString(w, "## Volumes\n")
	resp, err := c.etcd.KV.Range(c.getContext(), GetPrefix(MkKey("volumeid")))
	if err != nil {
		return err
	}
	for _, x := range resp.Kvs {
		io.WriteString(w, string(x.Key)+":\n")
		v := &models.Volume{}
		v.Unmarshal(x.Value)
		io.WriteString(w, v.String())
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "## INodes\n")
	resp, err = c.etcd.KV.Range(c.getContext(), GetPrefix(MkKey("volumemeta", "inode")))
	if err != nil {
		return err
	}
	for _, x := range resp.Kvs {
		io.WriteString(w, string(x.Key)+":\n")
		v := BytesToUint64(x.Value)
		io.WriteString(w, Uint64ToHex(v))
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "## BlockLocks\n")
	resp, err = c.etcd.KV.Range(c.getContext(), GetPrefix(MkKey("volumemeta", "blocklock")))
	if err != nil {
		return err
	}
	for _, x := range resp.Kvs {
		io.WriteString(w, string(x.Key)+":\n")
		io.WriteString(w, string(x.Value))
		io.WriteString(w, "\n")
	}
	return nil
}

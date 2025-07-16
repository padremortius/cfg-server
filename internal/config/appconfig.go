package config

type (
	App struct {
	}
)

func (c *Config) ReadPwd() error {
	pwd, err := fillPwdMap(c.SecPath)
	if err != nil {
		return err
	}

	c.Git.PrivateKey = pwd["git_PrivateKey"]
	c.Git.Password = pwd["git_password"]

	return nil
}

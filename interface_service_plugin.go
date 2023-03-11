/* Written by Dave Richards.
 * 
 * This is the top-level plugin interface, where you produce a new instance of the service.
 */
package discodove_interface_service

import (
	"log/syslog"
	"github.com/spf13/viper"
	"net"
	"github.com/mediumdaver/discodove-auth-plugin"
	"github.com/mediumdaver/discodove-user-lookup-plugin"
)
 
type DiscoDoveServiceFactory interface {
 
	/* This will be called once when we load this module a service instance, if you feel compelled to set something 
	 * up, perhaps a control/query/admin thread or something, then do it here in a controlled manner - similarly if 
	 * you want to pool connections, etc....  
	 *
	 * Each plugin is responsible for creating it's own syslog connection as *syslog.Writer has a mutex, and 
	 * I don't want the service threads to be blocking on writing to syslog - so you need to scale logging yourself.
	 * 
	 * We use Viper for config, and i will pass in the config directives for your module, but as it's viper you
	 * can access the entire discodove config too.  Feel free to specify your own config directives.
	 *
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove" - please
	 *            prefix your log messages with this and perhaps your own identifier e.g. "imap" or "pop3".
	 * syslogFacility : which facility to use in syslog.
	 * conf: a Viper subtree configuration for this service as specified in the discodove config.
	 * auth: a channel to send auth reqeusts to
	 * userlookup: a channel to send user lookup requests to
	 *
	 * The channels should be considered very scalable and should be re-used by many threads AND NOT CLOSED - please!
	 */
	Initialize(name string, syslogFacility syslog.Priority, auth chan discodove_interface_auth.DiscoDoveAuthRequest, userlookup chan discodove_interface_userlookup.DiscoDoveUserLookupRequest) error

	/* Because we can have a number of instances of the same service, we need seperate instances such that they may
	 * maintain seperate configurations. e.g. imap v's imaps, or different ports.  We call this function to produce a new instance
	 * of the service.
	 */
	NewService(conf *viper.Viper) DiscoDoveServicePlugin
}

type DiscoDoveServicePlugin interface {

	/* The real work is done here, this method will be called, in context of a go routine, for every 
	 * connection made on the listen port.  DiscoDove will not attempt to validate the connection in any
	 * way, the assumption is if it passed all the firewalls and discodove's controls (SSL for example),
	 * then it's over to you.
	 * 
	 * Please close the connection and return from this function when you're done.
	 *
	 * conn 			: the connection.
	 * alreadyEncrypted	: to let you know if the connection is already encrypted by discodove using TLS i.e. if it
	 *					  is, you probably do not want to let the client perform a STARTTLS
	 */
    HandleConnection(c net.Conn, alreadyEncrypted bool)
}